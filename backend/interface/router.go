package _interface

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/leo-yssa/monster-hunt-gamefi/backend/docs"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/infrastructure"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// --- Request Structs for Swagger ---
type RegisterPlayerRequest struct {
	Name string `json:"name"`
}

type AddMonsterRequest struct {
	Name   string `json:"name"`
	HP     int    `json:"hp"`
	Reward int    `json:"reward"`
}

type HuntMonsterRequest struct {
	MonsterID int64 `json:"monster_id"`
}

// 공통 핸들러 팩토리
func MakeTxQueueHandler(action string, paramBuilder func(*gin.Context) (map[string]interface{}, error), redisQueue *infrastructure.RedisQueue) gin.HandlerFunc {
	return func(c *gin.Context) {
		params, err := paramBuilder(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		txID := uuid.New().String()
		err = redisQueue.PushTxRequest(context.Background(), infrastructure.TxRequest{
			Action: action,
			Params: params,
			User:   txID,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"tx_id": txID, "status": "queued"})
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
		return map[string]interface{}{"monster_id": req.MonsterID}, nil
	}, redisQueue)
}

func NewRouter(redisQueue *infrastructure.RedisQueue) *gin.Engine {
	r := gin.Default()

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