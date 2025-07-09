package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/leo-yssa/monster-hunt-gamefi/backend/infrastructure"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"github.com/ethereum/go-ethereum/common"
)

type ConfirmationWorker struct {
	db     *gorm.DB
	client *ethclient.Client
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	shutdown chan struct{}
}

func NewConfirmationWorker() (*ConfirmationWorker, error) {
	db, err := infrastructure.InitGormDBFromEnv()
	if err != nil {
		return nil, err
	}
	
	rpcURL := infrastructure.LoadEnvOrDefault("MONSTER_GAME_RPC", "http://localhost:8545")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &ConfirmationWorker{
		db:       db,
		client:   client,
		ctx:      ctx,
		cancel:   cancel,
		shutdown: make(chan struct{}),
	}, nil
}

func (cw *ConfirmationWorker) Start() {
	log.Println("[CONFIRMATION WORKER] Pending 트랜잭션 상태 확인 시작...")

	// 시그널 핸들링
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 워커 고루틴 시작
	cw.wg.Add(1)
	go cw.processConfirmations()

	// 시그널 대기
	<-sigChan
	log.Println("[CONFIRMATION WORKER] 종료 신호 수신...")

	// 그레이스풀 종료
	cw.GracefulShutdown()
}

func (cw *ConfirmationWorker) processConfirmations() {
	defer cw.wg.Done()

	for {
		select {
		case <-cw.ctx.Done():
			log.Println("[CONFIRMATION WORKER] 컨텍스트 취소됨, 워커 종료")
			return
		case <-cw.shutdown:
			log.Println("[CONFIRMATION WORKER] 셧다운 신호 수신, 워커 종료")
			return
		default:
			// Pending 트랜잭션 조회
			var pendings []infrastructure.TxStatus
			if err := cw.db.WithContext(cw.ctx).Where("status = ?", "pending").Find(&pendings).Error; err != nil {
				log.Printf("[CONFIRMATION WORKER] DB 조회 에러: %v", err)
				time.Sleep(5 * time.Second)
				continue
			}

			// 각 트랜잭션 상태 확인 (고루틴으로 병렬 처리)
			for _, tx := range pendings {
				cw.wg.Add(1)
				go cw.checkAndUpdateTxStatus(&tx)
			}

			// 5초 대기
			time.Sleep(5 * time.Second)
		}
	}
}

func (cw *ConfirmationWorker) checkAndUpdateTxStatus(tx *infrastructure.TxStatus) {
	defer cw.wg.Done()

	// SafeExecute로 안전한 상태 확인
	err := infrastructure.SafeExecute(func() error {
		// RetryWithBackoff로 receipt 확인 안정성 확보
		return infrastructure.RetryWithBackoff(cw.ctx, func() error {
			// 컨텍스트 타임아웃 적용
			ctx, cancel := context.WithTimeout(cw.ctx, 10*time.Second)
			defer cancel()

			hash := common.HexToHash(tx.TxHash)
			receipt, err := cw.client.TransactionReceipt(ctx, hash)
			if err != nil {
				// receipt가 아직 없으면 skip (에러가 아님)
				return nil
			}

			// 상태 업데이트 (동시성 제어를 위한 조건부 업데이트)
			updateCtx, updateCancel := context.WithTimeout(cw.ctx, 5*time.Second)
			defer updateCancel()

			if receipt.Status == 1 {
				// pending 상태인 경우에만 success로 업데이트 (동시성 제어)
				result := cw.db.WithContext(updateCtx).Model(tx).
					Where("status = ?", "pending").
					Update("status", "success")
				
				if result.RowsAffected > 0 {
					log.Printf("[CONFIRMATION] tx %s success", tx.TxHash)
				}
			} else {
				// pending 상태인 경우에만 fail로 업데이트 (동시성 제어)
				result := cw.db.WithContext(updateCtx).Model(tx).
					Where("status = ?", "pending").
					Update("status", "fail")
				
				if result.RowsAffected > 0 {
					log.Printf("[CONFIRMATION] tx %s fail", tx.TxHash)
				}
			}

			return nil
		}, 2, 2*time.Second) // 최대 2회 재시도, 2초 초기 지연
	})

	if err != nil {
		log.Printf("[CONFIRMATION] SafeExecute 실패: %v", err)
	}
}

func (cw *ConfirmationWorker) GracefulShutdown() {
	log.Println("[CONFIRMATION WORKER] 그레이스풀 종료 시작...")

	// 컨텍스트 취소로 진행 중인 작업들 중단 신호
	cw.cancel()

	// 셧다운 신호 전송
	close(cw.shutdown)

	// 모든 고루틴 완료 대기 (30초 타임아웃)
	done := make(chan struct{})
	go func() {
		cw.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("[CONFIRMATION WORKER] 모든 작업 완료")
	case <-time.After(30 * time.Second):
		log.Println("[CONFIRMATION WORKER] 타임아웃으로 강제 종료")
	}

	// 리소스 정리
	if cw.client != nil {
		cw.client.Close()
	}

	if cw.db != nil {
		sqlDB, err := cw.db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}

	log.Println("[CONFIRMATION WORKER] 그레이스풀 종료 완료")
}

func main() {
	worker, err := NewConfirmationWorker()
	if err != nil {
		log.Fatalf("[CONFIRMATION WORKER] 워커 초기화 실패: %v", err)
	}

	worker.Start()
} 