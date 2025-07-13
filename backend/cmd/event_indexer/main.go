package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/config"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/domain"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/infrastructure"
)

type EventIndexer struct {
	repo   domain.EventRepository
	txStatusRepo domain.TxStatusRepository
	client *ethclient.Client
	ctx    context.Context
	cancel context.CancelFunc
	shutdown chan struct{}
}

type MonsterHuntedEvent struct {
	TxHash    string
	Player    string
	MonsterId string
	Reward    string
	CreatedAt time.Time
}

type PlayerRegisteredEvent struct {
	TxHash    string
	Player    string
	Name      string
	CreatedAt time.Time
}

func NewEventIndexer() (*EventIndexer, error) {
	db, err := config.InitGormDBFromEnv()
	if err != nil {
		return nil, err
	}

	repo := infrastructure.NewEventRepository(db)
	txStatusRepo := infrastructure.NewTxStatusRepository(db)

	rpcURL := config.LoadEnvOrDefault("MONSTER_GAME_WS_RPC", "ws://localhost:8545")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &EventIndexer{
		repo: repo,
		txStatusRepo: txStatusRepo,
		client: client,
		ctx: ctx,
		cancel: cancel,
		shutdown: make(chan struct{}),
	}, nil
}

func (ei *EventIndexer) Start() {
	log.Println("[EVENT INDEXER] 컨트랙트 이벤트 인덱싱 시작...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go ei.indexEvents()

	<-sigChan
	log.Println("[EVENT INDEXER] 종료 신호 수신...")
	ei.GracefulShutdown()
}

func (ei *EventIndexer) indexEvents() {
	contractAddr := os.Getenv("MONSTER_GAME_CONTRACT")
	if contractAddr == "" {
		log.Fatal("MONSTER_GAME_CONTRACT env not set")
	}
	address := common.HexToAddress(contractAddr)
	contract, err := infrastructure.NewContract(address, ei.client)
	if err != nil {
		log.Fatalf("컨트랙트 인스턴스 생성 실패: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	logs := make(chan types.Log)
	log.Println("[EVENT INDEXER] SubscribeFilterLogs 시작...")
	sub, err := ei.client.SubscribeFilterLogs(ei.ctx, query, logs)
	if err != nil {
		log.Fatalf("이벤트 구독 실패: %v", err)
	}
	log.Println("[EVENT INDEXER] SubscribeFilterLogs 성공, 이벤트 대기 중...")

	for {
		select {
		case <-ei.ctx.Done():
			log.Println("[EVENT INDEXER] 컨텍스트 취소됨, 인덱서 종료")
			return
		case <-ei.shutdown:
			log.Println("[EVENT INDEXER] 셧다운 신호 수신, 인덱서 종료")
			return
		case err := <-sub.Err():
			log.Printf("[EVENT INDEXER] Subscription error: %v", err)
			time.Sleep(5 * time.Second)
			continue
		case vLog := <-logs:
			log.Printf("[EVENT INDEXER] 이벤트 로그 수신: %+v", vLog)
			// MonsterHunted 이벤트 파싱
			if event, err := contract.ParseMonsterHunted(vLog); err == nil {
				log.Printf("[EVENT INDEXER] MonsterHunted 파싱 성공: %+v", event)
				ei.saveMonsterHuntedEvent(vLog.TxHash.Hex(), event)
				continue
			} else {
				log.Printf("[EVENT INDEXER] MonsterHunted 파싱 실패: %v", err)
			}
			// PlayerRegistered 이벤트 파싱
			if event, err := contract.ParsePlayerRegistered(vLog); err == nil {
				log.Printf("[EVENT INDEXER] PlayerRegistered 파싱 성공: %+v", event)
				ei.savePlayerRegisteredEvent(vLog.TxHash.Hex(), event)
				continue
			} else {
				log.Printf("[EVENT INDEXER] PlayerRegistered 파싱 실패: %v", err)
			}
		}
	}
}

func (ei *EventIndexer) saveMonsterHuntedEvent(txHash string, event *infrastructure.ContractMonsterHunted) {
	domainEvent := &domain.MonsterHuntedEvent{
		TxHash:    txHash,
		Player:    event.Player.Hex(),
		MonsterId: event.MonsterId.String(),
		Reward:    event.Reward.String(),
		CreatedAt: time.Now(),
	}
	if err := ei.repo.SaveMonsterHunted(ei.ctx, domainEvent); err != nil {
		log.Printf("[EVENT INDEXER] MonsterHunted 저장 실패: %v", err)
	}
	if err := ei.updateTxStatusSuccess(txHash); err != nil {
		log.Printf("[EVENT INDEXER] tx_status success 정정 실패: %v", err)
	}
}

func (ei *EventIndexer) savePlayerRegisteredEvent(txHash string, event *infrastructure.ContractPlayerRegistered) {
	domainEvent := &domain.PlayerRegisteredEvent{
		TxHash:    txHash,
		Player:    event.Player.Hex(),
		Name:      event.Name,
		CreatedAt: time.Now(),
	}
	if err := ei.repo.SavePlayerRegistered(ei.ctx, domainEvent); err != nil {
		log.Printf("[EVENT INDEXER] PlayerRegistered 저장 실패: %v", err)
	}
	if err := ei.updateTxStatusSuccess(txHash); err != nil {
		log.Printf("[EVENT INDEXER] tx_status success 정정 실패: %v", err)
	}
}

func (ei *EventIndexer) updateTxStatusSuccess(txHash string) error {
	return ei.txStatusRepo.UpdateTxStatus(ei.ctx, txHash, "success")
}

func (ei *EventIndexer) GracefulShutdown() {
	log.Println("[EVENT INDEXER] 그레이스풀 종료 시작...")
	ei.cancel()
	close(ei.shutdown)
	// DB/클라이언트 정리 (gorm은 자동, ethclient는 Close 필요)
	if ei.client != nil {
		ei.client.Close()
	}
	log.Println("[EVENT INDEXER] 그레이스풀 종료 완료")
}

func main() {
	_ = godotenv.Load()

	indexer, err := NewEventIndexer()
	if err != nil {
		log.Fatalf("EventIndexer 초기화 실패: %v", err)
	}

	indexer.Start()
} 