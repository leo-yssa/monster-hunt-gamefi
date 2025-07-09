package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Redis 큐 초기화
	redisQueue := infrastructure.NewRedisQueue()
	defer redisQueue.Close()

	// 라우터 설정
	r := _interface.NewRouter(redisQueue)

	// HTTP 서버 생성
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
		// 보안: 타임아웃 설정
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 그레이스풀 종료를 위한 채널
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 서버 시작 (고루틴에서)
	go func() {
		log.Println("[API] 서버 시작: http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[API] 서버 시작 실패: %v", err)
		}
	}()

	// 시그널 대기
	<-quit
	log.Println("[API] 서버 종료 신호 수신...")

	// 그레이스풀 종료 (30초 타임아웃)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("[API] 서버 강제 종료: %v", err)
	} else {
		log.Println("[API] 서버 그레이스풀 종료 완료")
	}
} 