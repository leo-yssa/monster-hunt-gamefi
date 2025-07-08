package main

import (
	"github.com/joho/godotenv"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/infrastructure"
	_interface "github.com/leo-yssa/monster-hunt-gamefi/backend/interface"
)

// @title Monster Hunt GameFi API
// @version 1.0
// @description Monster Hunt GameFi 백엔드 API 문서
// @host localhost:8080
// @BasePath /
func main() {
	_ = godotenv.Load()

	redisQueue := infrastructure.NewRedisQueue()
	r := _interface.NewRouter(redisQueue)
	r.Run(":8080")
} 