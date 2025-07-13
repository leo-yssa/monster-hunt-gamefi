package config

import (
	"os"
	"log"
	"time"
	"sync"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func LoadEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func InitRedisClientFromEnv() *redis.Client {
	redisURL := LoadEnvOrDefault("REDIS_URL", "redis://localhost:6379")
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Redis URL 파싱 실패: %v", err)
	}
	// Redis 연결 풀 최적화
	opt.PoolSize = 20
	opt.MinIdleConns = 5
	opt.PoolTimeout = 30 * time.Second
	opt.IdleTimeout = 5 * time.Minute
	opt.MaxRetries = 3
	opt.DialTimeout = 5 * time.Second
	opt.ReadTimeout = 3 * time.Second
	opt.WriteTimeout = 3 * time.Second
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
		dsn = "host=" + host +
			" port=" + port +
			" user=" + user +
			" password=" + password +
			" dbname=" + dbname +
			" sslmode=disable TimeZone=UTC" +
			" connect_timeout=10" +
			" statement_timeout=30000" +
			" idle_in_transaction_session_timeout=30000"
	}
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
		SkipDefaultTransaction: true,
	}
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(15 * time.Minute)
	return db, nil
}

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = InitGormDBFromEnv()
		if err != nil {
			log.Fatalf("DB 초기화 실패: %v", err)
		}
	})
	return dbInstance
} 