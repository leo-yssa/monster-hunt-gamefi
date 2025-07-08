package infrastructure

import (
	"os"
	"log"
	"github.com/go-redis/redis/v8"
)

func LoadEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func InitMonsterGameRepoFromEnv() (*MonsterGameRepository, error) {
	rpcURL := LoadEnvOrDefault("MONSTER_GAME_RPC", "http://localhost:8545")
	contractAddr := os.Getenv("MONSTER_GAME_CONTRACT")
	privKeyHex := os.Getenv("MONSTER_GAME_PRIVKEY")
	if contractAddr == "" || privKeyHex == "" {
		log.Fatal("컨트랙트 주소와 프라이빗키 환경변수(MONSTER_GAME_CONTRACT, MONSTER_GAME_PRIVKEY)가 필요합니다.")
	}
	return NewMonsterGameRepository(rpcURL, contractAddr, privKeyHex)
}

func InitRedisClientFromEnv() *redis.Client {
	redisURL := LoadEnvOrDefault("REDIS_URL", "redis://localhost:6379")
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Redis URL 파싱 실패: %v", err)
	}
	return redis.NewClient(opt)
} 