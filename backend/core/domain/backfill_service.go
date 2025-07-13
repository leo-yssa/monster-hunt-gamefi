package domain

import (
	"context"
	"time"
)

type BackfillResult struct {
	FromBlock    int64     `json:"from_block"`
	ToBlock      int64     `json:"to_block"`
	Inserted     int       `json:"inserted"`
	Skipped      int       `json:"skipped"`
	TotalLogs    int       `json:"total_logs"`
	Duration     string    `json:"duration"`
	CompletedAt  time.Time `json:"completed_at"`
}

type BackfillService interface {
	BackfillEvents(ctx context.Context, fromBlock, toBlock int64) (*BackfillResult, error)
} 