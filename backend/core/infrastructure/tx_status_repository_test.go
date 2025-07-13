package infrastructure

import (
	"context"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTxStatusTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&TxStatus{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestTxStatusRepositoryImpl_FindPendingTxs(t *testing.T) {
	db := setupTxStatusTestDB(t)
	repo := NewTxStatusRepository(db)
	// Insert test data
	db.Create(&TxStatus{TxHash: "0x1", Status: "pending", CreatedAt: time.Now()})
	db.Create(&TxStatus{TxHash: "0x2", Status: "success", CreatedAt: time.Now()})
	db.Create(&TxStatus{TxHash: "0x3", Status: "pending", CreatedAt: time.Now()})

	pendings, err := repo.FindPendingTxs(context.Background())
	if err != nil {
		t.Fatalf("FindPendingTxs failed: %v", err)
	}
	if len(pendings) != 2 {
		t.Errorf("expected 2 pending txs, got %d", len(pendings))
	}
}

func TestTxStatusRepositoryImpl_UpdateTxStatus(t *testing.T) {
	db := setupTxStatusTestDB(t)
	repo := NewTxStatusRepository(db)
	db.Create(&TxStatus{TxHash: "0x1", Status: "pending", CreatedAt: time.Now()})

	err := repo.UpdateTxStatus(context.Background(), "0x1", "success")
	if err != nil {
		t.Fatalf("UpdateTxStatus failed: %v", err)
	}
	var got TxStatus
	if err := db.First(&got, "tx_hash = ?", "0x1").Error; err != nil {
		t.Fatalf("tx not found: %v", err)
	}
	if got.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", got.Status)
	}
} 