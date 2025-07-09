package infrastructure

import (
	"context"
	"encoding/json"
	"os"
	"time"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisQueue struct {
	client    *redis.Client
	queueName string
}

type TxRequest struct {
	Action string                 `json:"action"`
	Params map[string]interface{} `json:"params"`
	User   string                 `json:"user"`
	Timestamp time.Time           `json:"timestamp"`
}

func NewRedisQueue() *RedisQueue {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379"
	}
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
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
	
	client := redis.NewClient(opt)
	
	// 연결 테스트
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Redis 연결 실패: %v", err)
	}
	
	return &RedisQueue{
		client:    client,
		queueName: "tx_queue",
	}
}

func (q *RedisQueue) PushTxRequest(ctx context.Context, req TxRequest) error {
	// 타임스탬프 추가
	if req.Timestamp.IsZero() {
		req.Timestamp = time.Now()
	}
	
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	
	// 파이프라인으로 성능 최적화
	pipe := q.client.Pipeline()
	pipe.RPush(ctx, q.queueName, data)
	pipe.Expire(ctx, q.queueName, 24*time.Hour) // 큐 만료 시간 설정
	
	_, err = pipe.Exec(ctx)
	return err
}

func (q *RedisQueue) PopTxRequest(ctx context.Context) (*TxRequest, error) {
	// BLPop으로 원자적 메시지 제거 (5초 타임아웃)
	result, err := q.client.BLPop(ctx, 5*time.Second, q.queueName).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // 큐가 비어있음
		}
		return nil, err
	}
	
	if len(result) < 2 {
		return nil, nil
	}
	
	var req TxRequest
	if err := json.Unmarshal([]byte(result[1]), &req); err != nil {
		return nil, err
	}
	
	return &req, nil
}

func (q *RedisQueue) GetQueueLength(ctx context.Context) (int64, error) {
	return q.client.LLen(ctx, q.queueName).Result()
}

func (q *RedisQueue) GetQueueStats(ctx context.Context) (map[string]interface{}, error) {
	length, err := q.GetQueueLength(ctx)
	if err != nil {
		return nil, err
	}
	
	// Redis 메모리 사용량
	info, err := q.client.Info(ctx, "memory").Result()
	if err != nil {
		return nil, err
	}
	
	return map[string]interface{}{
		"queue_length": length,
		"queue_name":   q.queueName,
		"redis_info":   info,
	}, nil
}

func (q *RedisQueue) Close() error {
	return q.client.Close()
}

// 성능 모니터링을 위한 헬퍼 함수들
func (q *RedisQueue) MonitorQueue(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			length, err := q.GetQueueLength(ctx)
			if err != nil {
				log.Printf("[REDIS] 큐 길이 조회 실패: %v", err)
				continue
			}
			
			if length > 100 {
				log.Printf("[REDIS] 큐 길이 경고: %d", length)
			}
		}
	}
} 