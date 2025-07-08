package main

import (
	"context"
	"log"
	"time"

	"github.com/leo-yssa/monster-hunt-gamefi/backend/infrastructure"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	db, err := infrastructure.InitGormDBFromEnv()
	if err != nil {
		log.Fatalf("Postgres DB 초기화 실패: %v", err)
	}
	rpcURL := infrastructure.LoadEnvOrDefault("MONSTER_GAME_RPC", "http://localhost:8545")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("ethclient 연결 실패: %v", err)
	}
	ctx := context.Background()

	log.Println("[CONFIRMATION WORKER] Pending 트랜잭션 상태 확인 시작...")
	for {
		var pendings []infrastructure.TxStatus
		if err := db.Where("status = ?", "pending").Find(&pendings).Error; err != nil {
			log.Printf("DB 조회 에러: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}
		for _, tx := range pendings {
			checkAndUpdateTxStatus(ctx, db, client, &tx)
		}
		time.Sleep(5 * time.Second)
	}
}

func checkAndUpdateTxStatus(ctx context.Context, db *gorm.DB, client *ethclient.Client, tx *infrastructure.TxStatus) {
	hash := common.HexToHash(tx.TxHash)
	receipt, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		// receipt가 아직 없으면 skip
		return
	}
	if receipt.Status == 1 {
		db.Model(tx).Update("status", "success")
		log.Printf("[CONFIRMATION] tx %s success", tx.TxHash)
	} else {
		db.Model(tx).Update("status", "fail")
		log.Printf("[CONFIRMATION] tx %s fail", tx.TxHash)
	}
} 