package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/config"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/application"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/infrastructure"
	_interface "github.com/leo-yssa/monster-hunt-gamefi/backend/core/interface"
)

func main() {
	_ = godotenv.Load()

	// Redis 큐 초기화
	redisQueue := infrastructure.NewRedisQueue()
	defer redisQueue.Close()

	// MonsterGameAdapter 초기화
	monsterGameAdapter, err := infrastructure.InitMonsterGameAdapterFromEnv()
	if err != nil {
		log.Fatalf("컨트랙트 연동 초기화 실패: %v", err)
	}

	// BackfillService 초기화
	backfillService, err := application.NewBackfillService(
		config.GetDB(),
		config.LoadEnvOrDefault("MONSTER_GAME_WS_RPC", "ws://localhost:8546"),
		os.Getenv("MONSTER_GAME_CONTRACT"),
	)
	if err != nil {
		log.Printf("Warning: Failed to initialize BackfillService: %v", err)
		backfillService = nil
	}

	// GameService 초기화
	gameService := application.NewGameService(monsterGameAdapter, backfillService)

	// PrometheusMonitor 초기화 및 모니터링 시작
	db := config.GetDB()
	monitor := infrastructure.NewPrometheusMonitor(db, redisQueue)
	go monitor.StartMonitoring(context.Background())

	// 라우터 설정
	r := _interface.NewRouter(redisQueue, gameService, monitor)

	// HTTP 서버 생성
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
		// 보안: 타임아웃 설정
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 서버 시작 (고루틴에서)
	go func() {
		log.Println("[API] 서버 시작: http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[API] 서버 시작 실패: %v", err)
		}
	}()

	// 그레이스풀 종료
	application.GracefulShutdownWithDefaultTimeout(srv)
} 