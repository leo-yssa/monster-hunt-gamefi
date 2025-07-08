package application

import (
	"github.com/leo-yssa/monster-hunt-gamefi/backend/infrastructure"
)

type GameService struct {
	MonsterGameRepo *infrastructure.MonsterGameRepository
}

func (s *GameService) RegisterPlayer(name string) (string, error) {
	return s.MonsterGameRepo.RegisterPlayer(name)
}

func (s *GameService) AddMonster(name string, hp, reward int) error {
	return s.MonsterGameRepo.AddMonster(name, hp, reward)
}

func (s *GameService) HuntMonster(monsterID int64) (string, error) {
	return s.MonsterGameRepo.HuntMonster(monsterID)
}
