package infrastructure

import (
	"time"
	"gorm.io/datatypes"
)

type TxStatus struct {
	ID        uint           `gorm:"primaryKey"`
	TxHash    string         `gorm:"size:66;uniqueIndex;not null"`
	Action    string         `gorm:"size:32;not null"`
	Params    datatypes.JSON `gorm:"type:jsonb"`
	UserID    string         `gorm:"size:64;index:idx_user_status_created,priority:1"`
	Status    string         `gorm:"size:16;not null;default:pending;index:idx_user_status_created,priority:2;index:idx_tx_status_status"`
	CreatedAt time.Time      `gorm:"index:idx_user_status_created,priority:3;index:idx_tx_status_created_at"`
	UpdatedAt time.Time
} 