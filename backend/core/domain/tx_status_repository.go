package domain

import (
	"context"
	"time"
)

type TxStatus struct {
	ID        uint
	TxHash    string
	Action    string
	Params    string // JSON 등으로 저장
	UserID    string
	Status    string
	SubmittedBlock uint64
	CreatedAt time.Time
}

type TxStatusRepository interface {
	FindPendingTxs(ctx context.Context) ([]TxStatus, error)
	UpdateTxStatus(ctx context.Context, txID string, status string) error
	CreateTxStatus(ctx context.Context, tx *TxStatus) error
} 