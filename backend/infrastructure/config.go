package infrastructure

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
	
	// Redis 연결 풀 최적화
	opt.PoolSize = 20                    // 기본 10에서 20으로 증가
	opt.MinIdleConns = 5                 // 최소 유휴 연결 수
	opt.PoolTimeout = 30 * time.Second   // 연결 풀 타임아웃
	opt.IdleTimeout = 5 * time.Minute    // 유휴 연결 타임아웃
	opt.MaxRetries = 3                   // 재시도 횟수
	opt.DialTimeout = 5 * time.Second    // 연결 타임아웃
	opt.ReadTimeout = 3 * time.Second    // 읽기 타임아웃
	opt.WriteTimeout = 3 * time.Second   // 쓰기 타임아웃
	
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
		
		// DB 연결 풀 최적화 파라미터 추가
		dsn = "host=" + host + 
			" port=" + port + 
			" user=" + user + 
			" password=" + password + 
			" dbname=" + dbname + 
			" sslmode=disable TimeZone=UTC" +
			" connect_timeout=10" +           // 연결 타임아웃
			" statement_timeout=30000" +      // 쿼리 타임아웃 (30초)
			" idle_in_transaction_session_timeout=30000" // 트랜잭션 타임아웃
	}
	
	// GORM 설정 최적화
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 프로덕션에서는 logger.Error로 변경
		PrepareStmt: true,                           // prepared statement 캐싱
		SkipDefaultTransaction: true,                 // 단일 쿼리에서 트랜잭션 스킵
	}
	
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	
	// DB 연결 풀 설정
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	
	// 연결 풀 최적화
	sqlDB.SetMaxIdleConns(10)     // 최대 유휴 연결 수
	sqlDB.SetMaxOpenConns(100)    // 최대 열린 연결 수
	sqlDB.SetConnMaxLifetime(time.Hour) // 연결 최대 수명
	sqlDB.SetConnMaxIdleTime(15 * time.Minute) // 유휴 연결 타임아웃
	
	if err := db.AutoMigrate(&TxStatus{}); err != nil {
		return nil, err
	}
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