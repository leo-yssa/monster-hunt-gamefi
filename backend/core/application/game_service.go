package application

import (
	"context"
	"fmt"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/domain"
)

// GameRepository 인터페이스 정의
type GameRepository interface {
	RegisterPlayer(ctx context.Context, name string) (string, error)
	AddMonster(ctx context.Context, name string, hp, reward int) (string, error)
	HuntMonster(ctx context.Context, monsterID int64) (string, error)
}

type GameService struct {
	monsterGamePort domain.MonsterGamePort
	backfillService domain.BackfillService
}

var _ domain.GameService = (*GameService)(nil)

func NewGameService(monsterGamePort domain.MonsterGamePort, backfillService domain.BackfillService) *GameService {
	return &GameService{
		monsterGamePort: monsterGamePort,
		backfillService: backfillService,
	}
}

func (s *GameService) RegisterPlayer(ctx context.Context, name string) (string, error) {
	return s.monsterGamePort.RegisterPlayer(ctx, name)
}

func (s *GameService) AddMonster(ctx context.Context, name string, hp, reward int) (string, error) {
	return s.monsterGamePort.AddMonster(ctx, name, hp, reward)
}

func (s *GameService) HuntMonster(ctx context.Context, monsterID int64) (string, error) {
	return s.monsterGamePort.HuntMonster(ctx, monsterID)
}

func (s *GameService) BackfillEvents(ctx context.Context, fromBlock, toBlock int64) (*domain.BackfillResult, error) {
	if s.backfillService == nil {
		return nil, fmt.Errorf("backfill service not available")
	}
	return s.backfillService.BackfillEvents(ctx, fromBlock, toBlock)
}
