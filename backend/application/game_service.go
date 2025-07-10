package application

import (
	"context"
)

// GameRepository 인터페이스 정의
type GameRepository interface {
	RegisterPlayer(ctx context.Context, name string) (string, error)
	AddMonster(ctx context.Context, name string, hp, reward int) (string, error)
	HuntMonster(ctx context.Context, monsterID int64) (string, error)
}

type GameService struct {
	repo GameRepository
}

func NewGameService(repo GameRepository) *GameService {
	return &GameService{
		repo: repo,
	}
}

func (s *GameService) RegisterPlayer(ctx context.Context, name string) (string, error) {
	return s.repo.RegisterPlayer(ctx, name)
}

func (s *GameService) AddMonster(ctx context.Context, name string, hp, reward int) (string, error) {
	return s.repo.AddMonster(ctx, name, hp, reward)
}

func (s *GameService) HuntMonster(ctx context.Context, monsterID int64) (string, error) {
	return s.repo.HuntMonster(ctx, monsterID)
}
