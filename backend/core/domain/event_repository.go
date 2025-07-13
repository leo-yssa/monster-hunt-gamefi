package domain

import (
	"context"
	"time"
)

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

type EventRepository interface {
	SaveMonsterHunted(ctx context.Context, event *MonsterHuntedEvent) error
	SavePlayerRegistered(ctx context.Context, event *PlayerRegisteredEvent) error
} 