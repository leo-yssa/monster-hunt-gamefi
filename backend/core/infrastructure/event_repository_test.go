package infrastructure

import (
	"context"
	"testing"
	"time"

	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&MonsterHuntedEvent{}, &PlayerRegisteredEvent{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestEventRepositoryImpl_SaveMonsterHunted(t *testing.T) {
	db := setupTestDB(t)
	repo := NewEventRepository(db)
	event := &domain.MonsterHuntedEvent{
		TxHash:    "0x123",
		Player:    "0xabc",
		MonsterId: "1",
		Reward:    "100",
		CreatedAt: time.Now(),
	}
	err := repo.SaveMonsterHunted(context.Background(), event)
	if err != nil {
		t.Fatalf("SaveMonsterHunted failed: %v", err)
	}
	var got MonsterHuntedEvent
	if err := db.First(&got, "tx_hash = ?", event.TxHash).Error; err != nil {
		t.Fatalf("event not found: %v", err)
	}
	if got.Player != event.Player || got.MonsterId != event.MonsterId {
		t.Errorf("saved event mismatch: got %+v, want %+v", got, event)
	}
}

func TestEventRepositoryImpl_SavePlayerRegistered(t *testing.T) {
	db := setupTestDB(t)
	repo := NewEventRepository(db)
	event := &domain.PlayerRegisteredEvent{
		TxHash:    "0x456",
		Player:    "0xdef",
		Name:      "Alice",
		CreatedAt: time.Now(),
	}
	err := repo.SavePlayerRegistered(context.Background(), event)
	if err != nil {
		t.Fatalf("SavePlayerRegistered failed: %v", err)
	}
	var got PlayerRegisteredEvent
	if err := db.First(&got, "tx_hash = ?", event.TxHash).Error; err != nil {
		t.Fatalf("event not found: %v", err)
	}
	if got.Player != event.Player || got.Name != event.Name {
		t.Errorf("saved event mismatch: got %+v, want %+v", got, event)
	}
} 