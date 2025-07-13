package domain

import "context"

type CurveLPStakingPort interface {
    Stake(ctx context.Context, user string, amount int64) (string, error)
    Unstake(ctx context.Context, user string, amount int64) (string, error)
    VotingPower(ctx context.Context, user string) (int64, error)
} 