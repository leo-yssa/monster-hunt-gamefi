package infrastructure

import (
	"context"
	"encoding/json"
	"os"

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
	client := redis.NewClient(opt)
	return &RedisQueue{
		client:    client,
		queueName: "tx_queue",
	}
}

func (q *RedisQueue) PushTxRequest(ctx context.Context, req TxRequest) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return q.client.RPush(ctx, q.queueName, data).Err()
} 