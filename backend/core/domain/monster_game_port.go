package domain

import "context"

type MonsterGamePort interface {
	RegisterPlayer(ctx context.Context, name string) (string, error)
	AddMonster(ctx context.Context, name string, hp, reward int) (string, error)
	HuntMonster(ctx context.Context, monsterID int64) (string, error)
} 