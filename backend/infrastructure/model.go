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
	UserID    string         `gorm:"size:64"`
	Status    string         `gorm:"size:16;not null;default:pending;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
} 