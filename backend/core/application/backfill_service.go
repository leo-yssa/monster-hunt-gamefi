package application

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/domain"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/infrastructure"
	"gorm.io/gorm"
)

// BackfillServiceImpl은 BackfillService의 구현체입니다.
type BackfillServiceImpl struct {
	db           *gorm.DB
	client       *ethclient.Client
	contract     *infrastructure.Contract
	eventRepo    domain.EventRepository
	wsURL        string
	contractAddr string
}

// NewBackfillService는 BackfillService의 새로운 인스턴스를 생성합니다.
func NewBackfillService(db *gorm.DB, wsURL, contractAddr string) (domain.BackfillService, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddr)
	contract, err := infrastructure.NewContract(address, client)
	if err != nil {
		return nil, err
	}

	eventRepo := infrastructure.NewEventRepository(db)

	return &BackfillServiceImpl{
		db:           db,
		client:       client,
		contract:     contract,
		eventRepo:    eventRepo,
		wsURL:        wsURL,
		contractAddr: contractAddr,
	}, nil
}

// BackfillEvents는 지정된 블록 범위의 이벤트를 복구합니다.
func (s *BackfillServiceImpl) BackfillEvents(ctx context.Context, fromBlock, toBlock int64) (*domain.BackfillResult, error) {
	startTime := time.Now()

	// 최신 블록 확인
	latest, err := s.client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	if toBlock == 0 || toBlock > int64(latest) {
		toBlock = int64(latest)
	}

	if fromBlock > toBlock {
		return nil, fmt.Errorf("from 블록이 to 블록보다 클 수 없습니다")
	}

	log.Printf("[BACKFILL] %d ~ %d 블록 범위 이벤트 인덱싱 시작...", fromBlock, toBlock)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{common.HexToAddress(s.contractAddr)},
	}

	logs, err := s.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	log.Printf("[BACKFILL] 총 %d개 로그 발견", len(logs))

	inserted, skipped := 0, 0

	for _, vLog := range logs {
		// MonsterHunted 이벤트 파싱
		if event, err := s.contract.ParseMonsterHunted(vLog); err == nil {
			record := &infrastructure.MonsterHuntedEvent{
				TxHash:    vLog.TxHash.Hex(),
				Player:    event.Player.Hex(),
				MonsterId: event.MonsterId.String(),
				Reward:    event.Reward.String(),
				CreatedAt: time.Now(),
			}
			
			// 중복 방지
			var cnt int64
			s.db.Model(&infrastructure.MonsterHuntedEvent{}).
				Where("tx_hash = ? AND player = ? AND monster_id = ?", record.TxHash, record.Player, record.MonsterId).
				Count(&cnt)
			
			if cnt == 0 {
				domainEvent := &domain.MonsterHuntedEvent{
					TxHash:    record.TxHash,
					Player:    record.Player,
					MonsterId: record.MonsterId,
					Reward:    record.Reward,
					CreatedAt: record.CreatedAt,
				}
				if err := s.eventRepo.SaveMonsterHunted(ctx, domainEvent); err != nil {
					log.Printf("[BACKFILL] MonsterHunted 저장 실패: %v", err)
				} else {
					inserted++
				}
			} else {
				skipped++
			}
			continue
		}

		// PlayerRegistered 이벤트 파싱
		if event, err := s.contract.ParsePlayerRegistered(vLog); err == nil {
			record := &infrastructure.PlayerRegisteredEvent{
				TxHash:    vLog.TxHash.Hex(),
				Player:    event.Player.Hex(),
				Name:      event.Name,
				CreatedAt: time.Now(),
			}
			
			var cnt int64
			s.db.Model(&infrastructure.PlayerRegisteredEvent{}).
				Where("tx_hash = ? AND player = ?", record.TxHash, record.Player).
				Count(&cnt)
			
			if cnt == 0 {
				domainEvent := &domain.PlayerRegisteredEvent{
					TxHash:    record.TxHash,
					Player:    record.Player,
					Name:      record.Name,
					CreatedAt: record.CreatedAt,
				}
				if err := s.eventRepo.SavePlayerRegistered(ctx, domainEvent); err != nil {
					log.Printf("[BACKFILL] PlayerRegistered 저장 실패: %v", err)
				} else {
					inserted++
				}
			} else {
				skipped++
			}
			continue
		}
	}

	duration := time.Since(startTime)
	
	result := &domain.BackfillResult{
		FromBlock:    fromBlock,
		ToBlock:      toBlock,
		Inserted:     inserted,
		Skipped:      skipped,
		TotalLogs:    len(logs),
		Duration:     duration.String(),
		CompletedAt:  time.Now(),
	}

	log.Printf("[BACKFILL] 저장: %d, 중복 스킵: %d, 소요시간: %s", inserted, skipped, duration)
	log.Println("[BACKFILL] 완료!")

	return result, nil
} 