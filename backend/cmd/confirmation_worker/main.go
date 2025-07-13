package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/leo-yssa/monster-hunt-gamefi/backend/config"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/domain"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/infrastructure"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/common"
)

type ConfirmationWorker struct {
	repo   domain.TxStatusRepository
	client *ethclient.Client
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	shutdown chan struct{}
}

func NewConfirmationWorker(repo domain.TxStatusRepository, client *ethclient.Client) *ConfirmationWorker {
	ctx, cancel := context.WithCancel(context.Background())
	return &ConfirmationWorker{
		repo:   repo,
		client: client,
		ctx:    ctx,
		cancel: cancel,
		shutdown: make(chan struct{}),
	}
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
			pendings, err := cw.repo.FindPendingTxs(cw.ctx)
			if err != nil {
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

func (cw *ConfirmationWorker) checkAndUpdateTxStatus(tx *domain.TxStatus) {
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
				err := cw.repo.UpdateTxStatus(updateCtx, tx.TxHash, "success")
				if err == nil {
					log.Printf("[CONFIRMATION] tx %s success", tx.TxHash)
				}
			} else {
				// pending 상태인 경우에만 fail로 업데이트 (동시성 제어)
				err := cw.repo.UpdateTxStatus(updateCtx, tx.TxHash, "fail")
				if err == nil {
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

const (
	PendingTimeout      = 30 * time.Minute
	PendingBlockWindow  = 20
)

func (cw *ConfirmationWorker) checkAndExpirePendingTxs() {
	now := time.Now()
	currentBlock, err := cw.client.BlockNumber(cw.ctx)
	if err != nil {
		log.Printf("[CONFIRMATION WORKER] 블록 넘버 조회 실패: %v", err)
		return
	}
	txs, err := cw.repo.FindPendingTxs(cw.ctx)
	if err != nil {
		log.Printf("[CONFIRMATION WORKER] pending 트랜잭션 조회 실패: %v", err)
		return
	}
	for _, tx := range txs {
		// created_at 기준 시간 초과
		if now.Sub(tx.CreatedAt) > PendingTimeout {
			cw.expireTx(&tx, "timeout")
			continue
		}
		// 블록 기준 초과 (tx_status에 submitted_block 필드가 있다고 가정)
		if tx.SubmittedBlock > 0 && currentBlock > tx.SubmittedBlock+PendingBlockWindow {
			cw.expireTx(&tx, "block_window")
		}
	}
}

func (cw *ConfirmationWorker) expireTx(tx *domain.TxStatus, reason string) {
	tx.Status = "expired"
	if err := cw.repo.UpdateTxStatus(context.Background(), tx.TxHash, "expired"); err != nil {
		log.Printf("[CONFIRMATION WORKER] 트랜잭션 만료 처리 실패: %v", err)
	} else {
		log.Printf("[CONFIRMATION WORKER] 트랜잭션 만료 처리: %s, reason=%s", tx.TxHash, reason)
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

	// DB 리소스는 프로그램 종료 시 자동으로 해제되므로 여기서는 별도 처리 안 함

	log.Println("[CONFIRMATION WORKER] 그레이스풀 종료 완료")
}

func main() {
	_ = godotenv.Load()
	db, err := config.InitGormDBFromEnv()
	if err != nil {
		log.Fatalf("DB 초기화 실패: %v", err)
	}
	repo := infrastructure.NewTxStatusRepository(db)

	rpcURL := config.LoadEnvOrDefault("MONSTER_GAME_RPC", "http://localhost:8545")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("이더리움 클라이언트 연결 실패: %v", err)
	}

	worker := NewConfirmationWorker(repo, client)
	worker.Start()
} 