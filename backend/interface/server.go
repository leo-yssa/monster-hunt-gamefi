package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/application"
	_ "github.com/leo-yssa/monster-hunt-gamefi/backend/docs"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/infrastructure"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Monster Hunt GameFi API
// @version 1.0
// @description Monster Hunt GameFi 백엔드 API 문서
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	playerRepo := infrastructure.NewInMemoryPlayerRepository()
	monsterRepo := infrastructure.NewInMemoryMonsterRepository()
	gameService := &application.GameService{
		PlayerRepo:  playerRepo,
		MonsterRepo: monsterRepo,
	}

	// @Summary 플레이어 등록
	// @Description address와 name으로 플레이어 등록
	// @Accept json
	// @Produce json
	// @Param player body struct{Address string; Name string} true "플레이어 정보"
	// @Success 200 {object} map[string]string
	// @Router /players [post]
	r.POST("/players", func(c *gin.Context) {
		var req struct {
			Address string `json:"address"`
			Name    string `json:"name"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := gameService.RegisterPlayer(req.Address, req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "registered"})
	})

	// @Summary 몬스터 추가
	// @Description name, hp, reward로 몬스터 추가
	// @Accept json
	// @Produce json
	// @Param monster body struct{Name string; HP int; Reward int} true "몬스터 정보"
	// @Success 200 {object} map[string]string
	// @Router /monsters [post]
	r.POST("/monsters", func(c *gin.Context) {
		var req struct {
			Name   string `json:"name"`
			HP     int    `json:"hp"`
			Reward int    `json:"reward"`
		}
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
	})

	// @Summary 사냥
	// @Description address와 monster_id로 사냥
	// @Accept json
	// @Produce json
	// @Param hunt body struct{Address string; MonsterID int} true "사냥 정보"
	// @Success 200 {object} map[string]int
	// @Router /hunt [post]
	r.POST("/hunt", func(c *gin.Context) {
		var req struct {
			Address   string `json:"address"`
			MonsterID int    `json:"monster_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reward, err := gameService.HuntMonster(req.Address, req.MonsterID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"reward": reward})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
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