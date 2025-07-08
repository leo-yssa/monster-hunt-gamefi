package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/application"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/infrastructure"
	_interface "github.com/leo-yssa/monster-hunt-gamefi/backend/interface"
)

// @title Monster Hunt GameFi API
// @version 1.0
// @description Monster Hunt GameFi 백엔드 API 문서
// @host localhost:8080
// @BasePath /
func main() {
	// .env 파일 자동 로드 (없어도 에러 무시)
	_ = godotenv.Load()

	// 환경변수 또는 config 패키지에서 읽기
	rpcURL := os.Getenv("MONSTER_GAME_RPC")
	if rpcURL == "" {
		rpcURL = "http://localhost:8545"
	}
	contractAddr := os.Getenv("MONSTER_GAME_CONTRACT")
	privKeyHex := os.Getenv("MONSTER_GAME_PRIVKEY")

	if contractAddr == "" || privKeyHex == "" {
		log.Fatal("컨트랙트 주소와 프라이빗키 환경변수(MONSTER_GAME_CONTRACT, MONSTER_GAME_PRIVKEY)가 필요합니다.")
	}

	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		log.Fatalf("프라이빗키 파싱 실패: %v", err)
	}
	publicAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	fmt.Printf("[INFO] 서버가 사용하는 계정 주소: %s\n", publicAddr.Hex())

	monsterGameRepo, err := infrastructure.NewMonsterGameRepository(rpcURL, contractAddr, privKeyHex)
	if err != nil {
		log.Fatalf("컨트랙트 연동 초기화 실패: %v", err)
	}
	gameService := &application.GameService{
		MonsterGameRepo: monsterGameRepo,
	}

	r := _interface.NewRouter(gameService)
	r.Run(":8080")
} 