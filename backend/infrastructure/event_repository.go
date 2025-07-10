package infrastructure

import "gorm.io/gorm"

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) SaveMonsterHunted(e *MonsterHuntedEvent) error {
	return r.db.Create(e).Error
}

func (r *EventRepository) SavePlayerRegistered(e *PlayerRegisteredEvent) error {
	return r.db.Create(e).Error
} 