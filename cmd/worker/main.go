package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/infrastructure"
)

type TxRequest struct {
	Action string                 `json:"action"` // "register", "hunt", "addMonster" 등
	Params map[string]interface{} `json:"params"`
	User   string                 `json:"user"`
}

func main() {
	rdb := infrastructure.InitRedisClientFromEnv()
	ctx := context.Background()
	queueName := "tx_queue"

	monsterGameRepo, err := infrastructure.InitMonsterGameRepoFromEnv()
	if err != nil {
		log.Fatalf("컨트랙트 연동 초기화 실패: %v", err)
	}

	log.Println("[WORKER] Redis 큐에서 트랜잭션 요청을 처리합니다...")
	for {
		res, err := rdb.BLPop(ctx, 5*time.Second, queueName).Result()
		if err == redis.Nil {
			continue // 큐에 데이터 없음
		} else if err != nil {
			log.Printf("Redis BLPop 에러: %v", err)
			continue
		}
		if len(res) < 2 {
			continue
		}
		var req TxRequest
		if err := json.Unmarshal([]byte(res[1]), &req); err != nil {
			log.Printf("TxRequest 파싱 실패: %v", err)
			continue
		}
		log.Printf("[WORKER] 트랜잭션 요청 처리: %+v", req)
		processTxRequest(ctx, monsterGameRepo, req)
	}
}

func processTxRequest(ctx context.Context, repo *infrastructure.MonsterGameRepository, req TxRequest) {
	switch req.Action {
	case "register":
		name, _ := req.Params["name"].(string)
		txHash, err := repo.RegisterPlayer(name)
		if err != nil {
			log.Printf("registerPlayer 실패: %v", err)
		} else {
			log.Printf("registerPlayer tx: %s", txHash)
		}
	case "addMonster":
		name, _ := req.Params["name"].(string)
		hp, _ := toInt(req.Params["hp"])
		reward, _ := toInt(req.Params["reward"])
		err := repo.AddMonster(name, hp, reward)
		if err != nil {
			log.Printf("addMonster 실패: %v", err)
		} else {
			log.Printf("addMonster 성공: %s", name)
		}
	case "hunt":
		monsterID, _ := toInt64(req.Params["monster_id"])
		txHash, err := repo.HuntMonster(monsterID)
		if err != nil {
			log.Printf("huntMonster 실패: %v", err)
		} else {
			log.Printf("huntMonster tx: %s", txHash)
		}
	default:
		log.Printf("알 수 없는 액션: %s", req.Action)
	}
}

func toInt(v interface{}) (int, bool) {
	switch val := v.(type) {
	case float64:
		return int(val), true
	case int:
		return val, true
	case int64:
		return int(val), true
	case string:
		var i int
		_, err := fmt.Sscanf(val, "%d", &i)
		return i, err == nil
	default:
		return 0, false
	}
}

func toInt64(v interface{}) (int64, bool) {
	switch val := v.(type) {
	case float64:
		return int64(val), true
	case int:
		return int64(val), true
	case int64:
		return val, true
	case string:
		var i int64
		_, err := fmt.Sscanf(val, "%d", &i)
		return i, err == nil
	default:
		return 0, false
	}
} 