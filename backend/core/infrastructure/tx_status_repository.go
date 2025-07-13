package infrastructure

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/datatypes"

	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/domain"
)

type TxStatusRepositoryImpl struct {
	db *gorm.DB
}

func NewTxStatusRepository(db *gorm.DB) *TxStatusRepositoryImpl {
	return &TxStatusRepositoryImpl{db: db}
}

// FindPendingTxs는 status가 'pending'인 트랜잭션을 모두 반환합니다.
func (r *TxStatusRepositoryImpl) FindPendingTxs(ctx context.Context) ([]domain.TxStatus, error) {
	var models []TxStatus
	if err := r.db.WithContext(ctx).Where("status = ?", "pending").Find(&models).Error; err != nil {
		return nil, err
	}
	var result []domain.TxStatus
	for _, m := range models {
		result = append(result, toDomainTxStatus(m))
	}
	return result, nil
}

// UpdateTxStatus는 트랜잭션 상태를 갱신합니다.
func (r *TxStatusRepositoryImpl) UpdateTxStatus(ctx context.Context, txID string, status string) error {
	return r.db.WithContext(ctx).Model(&TxStatus{}).Where("tx_hash = ?", txID).Update("status", status).Error
}

func (r *TxStatusRepositoryImpl) CreateTxStatus(ctx context.Context, tx *domain.TxStatus) error {
	record := toInfraTxStatus(tx)
	return r.db.WithContext(ctx).Create(&record).Error
}

func toDomainTxStatus(m TxStatus) domain.TxStatus {
	return domain.TxStatus{
		ID:        m.ID,
		TxHash:    m.TxHash,
		Action:    m.Action,
		Params:    string(m.Params),
		UserID:    m.UserID,
		Status:    m.Status,
		SubmittedBlock: m.SubmittedBlock,
		CreatedAt: m.CreatedAt,
	}
}

func toInfraTxStatus(tx *domain.TxStatus) TxStatus {
	return TxStatus{
		TxHash:    tx.TxHash,
		Action:    tx.Action,
		Params:    datatypes.JSON([]byte(tx.Params)),
		UserID:    tx.UserID,
		Status:    tx.Status,
		SubmittedBlock: tx.SubmittedBlock,
		CreatedAt: tx.CreatedAt,
	}
} 