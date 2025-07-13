package _interface

import (
	"context"
	"log"
	"net/http"
	"time"

	"bytes"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/application"
	_ "github.com/leo-yssa/monster-hunt-gamefi/backend/core/docs"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/infrastructure"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

// --- Request Structs for Swagger ---
type RegisterPlayerRequest struct {
	Name string `json:"name" binding:"required,min=1,max=32"`
}

type AddMonsterRequest struct {
	Name   string `json:"name" binding:"required,min=1,max=50"`
	HP     int    `json:"hp" binding:"required,min=1,max=10000"`
	Reward int    `json:"reward" binding:"required,min=1,max=1000000"`
}

type HuntMonsterRequest struct {
	MonsterID int `json:"monster_id" binding:"min=0"`
}

type BackfillRequest struct {
	FromBlock int64 `json:"from_block" binding:"min=0"`
	ToBlock   int64 `json:"to_block" binding:"min=0"`
}

type Handler struct {
	RedisQueue  *infrastructure.RedisQueue
	GameService *application.GameService
	Monitor     *infrastructure.PrometheusMonitor
}

// --- Curve LP Staking 관련 Structs ---
type StakeCurveLPRequest struct {
    User      string `json:"user"`
    Amount    int64  `json:"amount"`
    PrivKeyHex string `json:"privKeyHex"`
}

// 보안 미들웨어: 클라이언트 IP 로깅 및 검증
func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := infrastructure.GetClientIP(c.Request)
		
		// 프라이빗 IP 접근 제한 (선택적)
		if infrastructure.IsPrivateIP(clientIP) {
			// 개발 환경에서는 허용, 프로덕션에서는 제한 가능
			// c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Private IP access not allowed"})
			// return
		}

		// 클라이언트 IP를 컨텍스트에 저장
		c.Set("client_ip", clientIP)
		
		// 요청 로깅
		log.Printf("[SECURITY] Request from IP: %s, Path: %s, Method: %s", 
			clientIP, c.Request.URL.Path, c.Request.Method)
		
		c.Next()
	}
}

// @Summary 플레이어 등록
// @Description address와 name으로 플레이어 등록
// @Accept json
// @Produce json
// @Param player body RegisterPlayerRequest true "플레이어 정보"
// @Success 200 {object} map[string]string
// @Router /players [post]
func (h *Handler) RegisterPlayerHandler(c *gin.Context) {
	start := time.Now()
	var req RegisterPlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Monitor.RecordTransaction("fail", time.Since(start))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(req.Name) == 0 || len(req.Name) > 32 {
		h.Monitor.RecordTransaction("fail", time.Since(start))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid name length"})
		return
	}
	txID := uuid.New().String()
	err := h.RedisQueue.PushTxRequest(context.Background(), infrastructure.TxRequest{
		Action: "register",
		Params: map[string]interface{}{ "name": req.Name },
		User:   txID,
	})
	if err != nil {
		h.Monitor.RecordTransaction("fail", time.Since(start))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.Monitor.RecordTransaction("success", time.Since(start))
	c.JSON(http.StatusOK, gin.H{"tx_id": txID, "status": "queued"})
}

// @Summary 몬스터 추가
// @Description name, hp, reward로 몬스터 추가
// @Accept json
// @Produce json
// @Param monster body AddMonsterRequest true "몬스터 정보"
// @Success 200 {object} map[string]string
// @Router /monsters [post]
func (h *Handler) AddMonsterHandler(c *gin.Context) {
	start := time.Now()
	var req AddMonsterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Monitor.RecordTransaction("fail", time.Since(start))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.HP <= 0 || req.HP > 10000 {
		h.Monitor.RecordTransaction("fail", time.Since(start))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid HP value"})
		return
	}
	if req.Reward <= 0 || req.Reward > 1000000 {
		h.Monitor.RecordTransaction("fail", time.Since(start))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reward value"})
		return
	}
	txID := uuid.New().String()
	err := h.RedisQueue.PushTxRequest(context.Background(), infrastructure.TxRequest{
		Action: "addMonster",
		Params: map[string]interface{}{ "name": req.Name, "hp": req.HP, "reward": req.Reward },
		User:   txID,
	})
	if err != nil {
		h.Monitor.RecordTransaction("fail", time.Since(start))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.Monitor.RecordTransaction("success", time.Since(start))
	c.JSON(http.StatusOK, gin.H{"tx_id": txID, "status": "queued"})
}

// @Summary 사냥
// @Description address와 monster_id로 사냥
// @Accept json
// @Produce json
// @Param hunt body HuntMonsterRequest true "사냥 정보"
// @Success 200 {object} map[string]int
// @Router /hunt [post]
func (h *Handler) HuntMonsterHandler(c *gin.Context) {
	start := time.Now()
	body, _ := io.ReadAll(c.Request.Body)
	log.Printf("[DEBUG] RAW BODY: %s", string(body))
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	var req HuntMonsterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Monitor.RecordTransaction("fail", time.Since(start))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.MonsterID < 0 {
		h.Monitor.RecordTransaction("fail", time.Since(start))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid monster ID"})
		return
	}
	txID := uuid.New().String()
	err := h.RedisQueue.PushTxRequest(context.Background(), infrastructure.TxRequest{
		Action: "hunt",
		Params: map[string]interface{}{ "monster_id": req.MonsterID },
		User:   txID,
	})
	if err != nil {
		h.Monitor.RecordTransaction("fail", time.Since(start))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.Monitor.RecordTransaction("success", time.Since(start))
	c.JSON(http.StatusOK, gin.H{"tx_id": txID, "status": "queued"})
}

// @Summary 이벤트 백필드
// @Description 지정된 블록 범위의 이벤트를 백필드로 인덱싱
// @Accept json
// @Produce json
// @Param backfill body BackfillRequest true "백필드 정보"
// @Success 200 {object} domain.BackfillResult
// @Router /admin/backfill [post]
func (h *Handler) BackfillHandler(c *gin.Context) {
	var req BackfillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.FromBlock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from_block must be non-negative"})
		return
	}
	if req.ToBlock > 0 && req.FromBlock > req.ToBlock {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from_block cannot be greater than to_block"})
		return
	}
	result, err := h.GameService.BackfillEvents(c.Request.Context(), req.FromBlock, req.ToBlock)
	if err != nil {
		log.Printf("[BACKFILL API] Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Curve LP 토큰 스테이킹
// @Description Curve LP 토큰을 스테이킹합니다. (privKeyHex는 유저의 개인키이며, 실제 서비스에서는 절대 서버로 전달하면 안 됩니다. 테스트/개발용)
// @Accept json
// @Produce json
// @Param stake body StakeCurveLPRequest true "스테이킹 정보 (user: 유저 주소, amount: 수량, privKeyHex: 유저 개인키)"
// @Success 200 {object} map[string]string
// @Router /stake/curve-lp [post]
func (h *Handler) StakeCurveLPHandler(c *gin.Context) {
    var req StakeCurveLPRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    txID := uuid.New().String()
    err := h.RedisQueue.PushTxRequest(context.Background(), infrastructure.TxRequest{
        Action: "stakeCurveLP",
        Params: map[string]interface{}{ "amount": req.Amount, "privKeyHex": req.PrivKeyHex },
        User:   txID,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"tx_id": txID, "status": "queued"})
}

// @Summary Curve LP 토큰 언스테이킹
// @Description Curve LP 토큰을 언스테이킹합니다. (privKeyHex는 유저의 개인키이며, 실제 서비스에서는 절대 서버로 전달하면 안 됩니다. 테스트/개발용)
// @Accept json
// @Produce json
// @Param unstake body StakeCurveLPRequest true "언스테이킹 정보 (user: 유저 주소, amount: 수량, privKeyHex: 유저 개인키)"
// @Success 200 {object} map[string]string
// @Router /unstake/curve-lp [post]
func (h *Handler) UnstakeCurveLPHandler(c *gin.Context) {
    var req StakeCurveLPRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    txID := uuid.New().String()
    err := h.RedisQueue.PushTxRequest(context.Background(), infrastructure.TxRequest{
        Action: "unstakeCurveLP",
        Params: map[string]interface{}{ "amount": req.Amount, "privKeyHex": req.PrivKeyHex },
        User:   txID,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"tx_id": txID, "status": "queued"})
}

// @Summary Curve LP 스테이킹 현황 조회
// @Description 유저의 Curve LP 스테이킹 현황을 조회합니다.
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /stake/curve-lp [get]
func (h *Handler) GetCurveLPStakeHandler(c *gin.Context) {
    // TODO: 실제 스테이킹 현황 조회 로직 연결
    c.JSON(http.StatusOK, gin.H{"staked": 0, "voting_power": 0})
}

// @Summary 프로메테우스 메트릭
// @Description 프로메테우스 형식의 메트릭 노출
// @Accept json
// @Produce text/plain
// @Success 200 {string} string
// @Router /metrics [get]
func MetricsHandler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}

func NewRouter(redisQueue *infrastructure.RedisQueue, gameService *application.GameService, monitor *infrastructure.PrometheusMonitor) *gin.Engine {
	h := &Handler{RedisQueue: redisQueue, GameService: gameService, Monitor: monitor}
	r := gin.Default()

	r.Use(SecurityMiddleware())

	if os.Getenv("RATE_LIMIT_ENABLED") != "false" {
		rate := os.Getenv("RATE_LIMIT_RATE")
		if rate == "" {
			rate = "1000-M"
		}
		limiter := limiter.New(memory.NewStore(), limiter.Rate{
			Period: 1 * time.Minute,
			Limit:  1000,
		})
		r.Use(ginlimiter.NewMiddleware(limiter))
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "timestamp": time.Now().Unix()})
	})
	
	r.GET("/metrics", MetricsHandler())

	r.POST("/players", h.RegisterPlayerHandler)
	r.POST("/monsters", h.AddMonsterHandler)
	r.POST("/hunt", h.HuntMonsterHandler)
	r.POST("/admin/backfill", h.BackfillHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Curve LP Staking 관련 엔드포인트
	r.POST("/stake/curve-lp", h.StakeCurveLPHandler)
	r.POST("/unstake/curve-lp", h.UnstakeCurveLPHandler)
	r.GET("/stake/curve-lp", h.GetCurveLPStakeHandler)

	return r
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing token"})
			return
		}
		address, err := infrastructure.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}
		c.Set("address", address)
		c.Next()
	}
} 