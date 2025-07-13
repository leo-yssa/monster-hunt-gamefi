package infrastructure

import (
	"context"
	"gorm.io/gorm"

	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/domain"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db: db}
}

func (r *EventRepositoryImpl) SaveMonsterHunted(ctx context.Context, event *domain.MonsterHuntedEvent) error {
	record := &MonsterHuntedEvent{
		TxHash:    event.TxHash,
		Player:    event.Player,
		MonsterId: event.MonsterId,
		Reward:    event.Reward,
		CreatedAt: event.CreatedAt,
	}
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *EventRepositoryImpl) SavePlayerRegistered(ctx context.Context, event *domain.PlayerRegisteredEvent) error {
	record := &PlayerRegisteredEvent{
		TxHash:    event.TxHash,
		Player:    event.Player,
		Name:      event.Name,
		CreatedAt: event.CreatedAt,
	}
	return r.db.WithContext(ctx).Create(record).Error
}