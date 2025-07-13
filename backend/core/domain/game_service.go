package domain

import (
	"context"
)

// GameService는 게임 도메인 서비스의 인터페이스입니다.
type GameService interface {
	RegisterPlayer(ctx context.Context, name string) (string, error)
	AddMonster(ctx context.Context, name string, hp, reward int) (string, error)
	HuntMonster(ctx context.Context, monsterID int64) (string, error)
	BackfillEvents(ctx context.Context, fromBlock, toBlock int64) (*BackfillResult, error)
} 