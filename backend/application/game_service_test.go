package application

import (
	"testing"

	"github.com/leo-yssa/monster-hunt-gamefi/backend/infrastructure"
)

func TestRegisterPlayer(t *testing.T) {
	playerRepo := infrastructure.NewInMemoryPlayerRepository()
	monsterRepo := infrastructure.NewInMemoryMonsterRepository()
	service := &GameService{PlayerRepo: playerRepo, MonsterRepo: monsterRepo}

	err := service.RegisterPlayer("0x123", "Alice")
	if err != nil {
		t.Fatal(err)
	}
	player, err := playerRepo.FindByAddress("0x123")
	if err != nil || player.Name != "Alice" {
		t.Errorf("player registration failed")
	}
} 