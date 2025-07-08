package infrastructure

import (
	"os"
	"log"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
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

func InitGormDBFromEnv() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		host := LoadEnvOrDefault("POSTGRES_HOST", "localhost")
		port := LoadEnvOrDefault("POSTGRES_PORT", "5432")
		user := LoadEnvOrDefault("POSTGRES_USER", "postgres")
		password := LoadEnvOrDefault("POSTGRES_PASSWORD", "postgres")
		dbname := LoadEnvOrDefault("POSTGRES_DB", "monster_gamefi")
		dsn = "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable TimeZone=UTC"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&TxStatus{}); err != nil {
		return nil, err
	}
	return db, nil
} 