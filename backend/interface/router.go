package _interface

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/leo-yssa/monster-hunt-gamefi/backend/docs"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/infrastructure"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	MonsterID int64 `json:"monster_id" binding:"required,min=0"`
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

// 공통 핸들러 팩토리 (보안 강화)
func MakeTxQueueHandler(action string, paramBuilder func(*gin.Context) (map[string]interface{}, error), redisQueue *infrastructure.RedisQueue) gin.HandlerFunc {
	return func(c *gin.Context) {
		// SafeExecute로 안전한 실행
		err := infrastructure.SafeExecute(func() error {
			params, err := paramBuilder(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return err
			}
			
			txID := uuid.New().String()
			err = redisQueue.PushTxRequest(context.Background(), infrastructure.TxRequest{
				Action: action,
				Params: params,
				User:   txID,
			})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return err
			}
			
			c.JSON(http.StatusOK, gin.H{"tx_id": txID, "status": "queued"})
			return nil
		})
		
		if err != nil {
			log.Printf("[ERROR] Handler execution failed: %v", err)
			// SafeExecute에서 이미 응답을 보냈으므로 여기서는 로깅만
		}
	}
}

// @Summary 플레이어 등록
// @Description address와 name으로 플레이어 등록
// @Accept json
// @Produce json
// @Param player body RegisterPlayerRequest true "플레이어 정보"
// @Success 200 {object} map[string]string
// @Router /players [post]
func RegisterPlayerHandler(redisQueue *infrastructure.RedisQueue) gin.HandlerFunc {
	return MakeTxQueueHandler("register", func(c *gin.Context) (map[string]interface{}, error) {
		var req RegisterPlayerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			return nil, err
		}
		
		// 추가 입력 검증
		if len(req.Name) == 0 || len(req.Name) > 32 {
			return nil, fmt.Errorf("invalid name length")
		}
		
		return map[string]interface{}{"name": req.Name}, nil
	}, redisQueue)
}

// @Summary 몬스터 추가
// @Description name, hp, reward로 몬스터 추가
// @Accept json
// @Produce json
// @Param monster body AddMonsterRequest true "몬스터 정보"
// @Success 200 {object} map[string]string
// @Router /monsters [post]
func AddMonsterHandler(redisQueue *infrastructure.RedisQueue) gin.HandlerFunc {
	return MakeTxQueueHandler("addMonster", func(c *gin.Context) (map[string]interface{}, error) {
		var req AddMonsterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			return nil, err
		}
		
		// 추가 입력 검증
		if req.HP <= 0 || req.HP > 10000 {
			return nil, fmt.Errorf("invalid HP value")
		}
		if req.Reward <= 0 || req.Reward > 1000000 {
			return nil, fmt.Errorf("invalid reward value")
		}
		
		return map[string]interface{}{"name": req.Name, "hp": req.HP, "reward": req.Reward}, nil
	}, redisQueue)
}

// @Summary 사냥
// @Description address와 monster_id로 사냥
// @Accept json
// @Produce json
// @Param hunt body HuntMonsterRequest true "사냥 정보"
// @Success 200 {object} map[string]int
// @Router /hunt [post]
func HuntMonsterHandler(redisQueue *infrastructure.RedisQueue) gin.HandlerFunc {
	return MakeTxQueueHandler("hunt", func(c *gin.Context) (map[string]interface{}, error) {
		var req HuntMonsterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			return nil, err
		}
		
		// 추가 입력 검증
		if req.MonsterID < 0 {
			return nil, fmt.Errorf("invalid monster ID")
		}
		
		return map[string]interface{}{"monster_id": req.MonsterID}, nil
	}, redisQueue)
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

func NewRouter(redisQueue *infrastructure.RedisQueue) *gin.Engine {
	r := gin.Default()
	
	// 보안 미들웨어 추가
	r.Use(SecurityMiddleware())
	
	// Rate Limiting 설정
	if os.Getenv("RATE_LIMIT_ENABLED") != "false" {
		rate := os.Getenv("RATE_LIMIT_RATE")
		if rate == "" {
			rate = "1000-M" // 기본값: 1분 1000회
		}
		
		limiter := limiter.New(memory.NewStore(), limiter.Rate{
			Period: 1 * time.Minute,
			Limit:  1000,
		})
		
		r.Use(ginlimiter.NewMiddleware(limiter))
	}
	
	// 헬스체크 엔드포인트
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "timestamp": time.Now().Unix()})
	})
	
	// 프로메테우스 메트릭 엔드포인트
	r.GET("/metrics", MetricsHandler())
	
	// API 엔드포인트
	r.POST("/players", RegisterPlayerHandler(redisQueue))
	r.POST("/monsters", AddMonsterHandler(redisQueue))
	r.POST("/hunt", HuntMonsterHandler(redisQueue))
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
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