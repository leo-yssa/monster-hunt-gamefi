package domain

import "context"

type CurveLPTokenPort interface {
    Approve(ctx context.Context, owner, spender string, amount int64) (string, error)
    BalanceOf(ctx context.Context, owner string) (int64, error)
} 