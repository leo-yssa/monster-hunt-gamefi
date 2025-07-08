package infrastructure

import (
	"os"
	"testing"
)

func TestMonsterGameRepository_Integration(t *testing.T) {
	// 환경변수 또는 하드코딩 값 사용 (실제 배포 후 값으로 교체)
	rpcURL := "http://localhost:8545"
	contractAddr := os.Getenv("MONSTER_GAME_CONTRACT") // 또는 "0x..."
	privKeyHex := os.Getenv("MONSTER_GAME_PRIVKEY")   // 또는 "..."

	if contractAddr == "" || privKeyHex == "" {
		t.Skip("컨트랙트 주소와 프라이빗키 환경변수 필요")
	}

	repo, err := NewMonsterGameRepository(rpcURL, contractAddr, privKeyHex)
	if err != nil {
		t.Fatalf("MonsterGameRepository 생성 실패: %v", err)
	}

	// 1. RegisterPlayer
	playerName := "Alice"
	txHash, err := repo.RegisterPlayer(playerName)
	if err != nil {
		t.Fatalf("RegisterPlayer 실패: %v", err)
	}
	t.Logf("RegisterPlayer tx: %s", txHash)

	// 2. HuntMonster (monsterID=0 가정)
	txHash, err = repo.HuntMonster(0)
	if err != nil {
		t.Fatalf("HuntMonster 실패: %v", err)
	}
	t.Logf("HuntMonster tx: %s", txHash)
} 