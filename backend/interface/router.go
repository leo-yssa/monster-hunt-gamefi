package _interface

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/application"
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

// @Summary 플레이어 등록
// @Description address와 name으로 플레이어 등록
// @Accept json
// @Produce json
// @Param player body RegisterPlayerRequest true "플레이어 정보"
// @Success 200 {object} map[string]string
// @Router /players [post]
func RegisterPlayerHandler(gameService *application.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterPlayerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		txHash, err := gameService.RegisterPlayer(req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"tx_hash": txHash})
	}
}

// @Summary 몬스터 추가
// @Description name, hp, reward로 몬스터 추가
// @Accept json
// @Produce json
// @Param monster body AddMonsterRequest true "몬스터 정보"
// @Success 200 {object} map[string]string
// @Router /monsters [post]
func AddMonsterHandler(gameService *application.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddMonsterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := gameService.AddMonster(req.Name, req.HP, req.Reward)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "monster added"})
	}
}

// @Summary 사냥
// @Description address와 monster_id로 사냥
// @Accept json
// @Produce json
// @Param hunt body HuntMonsterRequest true "사냥 정보"
// @Success 200 {object} map[string]int
// @Router /hunt [post]
func HuntMonsterHandler(gameService *application.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req HuntMonsterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		txHash, err := gameService.HuntMonster(req.MonsterID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"tx_hash": txHash})
	}
}

func NewRouter(gameService *application.GameService) *gin.Engine {
	r := gin.Default()

	r.POST("/players", RegisterPlayerHandler(gameService))
	r.POST("/monsters", AddMonsterHandler(gameService))
	r.POST("/hunt", HuntMonsterHandler(gameService))

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