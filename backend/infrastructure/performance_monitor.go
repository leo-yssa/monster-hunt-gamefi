package infrastructure

import (
	"context"
	"log"
	"runtime"
	"time"
	"gorm.io/gorm"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Prometheus 메트릭 정의
var (
	// 시스템 메트릭
	goroutinesGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "app_goroutines_total",
		Help: "현재 고루틴 수",
	})
	
	memoryAllocGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "app_memory_alloc_bytes",
		Help: "현재 할당된 메모리 (bytes)",
	})
	
	memorySysGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "app_memory_sys_bytes", 
		Help: "시스템에서 사용 중인 메모리 (bytes)",
	})
	
	// DB 메트릭
	dbConnectionsGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "app_db_connections_total",
		Help: "DB 연결 수",
	})
	
	dbConnectionsInUseGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "app_db_connections_in_use",
		Help: "사용 중인 DB 연결 수",
	})
	
	dbConnectionsIdleGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "app_db_connections_idle",
		Help: "유휴 DB 연결 수",
	})
	
	// Redis 메트릭
	redisQueueLengthGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "app_redis_queue_length",
		Help: "Redis 큐 길이",
	})
	
	// 트랜잭션 메트릭
	transactionProcessingTimeHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "app_transaction_processing_seconds",
		Help: "트랜잭션 처리 시간 (초)",
		Buckets: prometheus.DefBuckets,
	})
	
	transactionStatusCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "app_transaction_status_total",
		Help: "트랜잭션 상태별 카운트",
	}, []string{"status"})
)

// PrometheusMonitor 프로메테우스 기반 모니터링
type PrometheusMonitor struct {
	db          *gorm.DB
	redisQueue  *RedisQueue
	startTime   time.Time
}

// NewPrometheusMonitor 새로운 프로메테우스 모니터 생성
func NewPrometheusMonitor(db *gorm.DB, redisQueue *RedisQueue) *PrometheusMonitor {
	return &PrometheusMonitor{
		db:         db,
		redisQueue: redisQueue,
		startTime:  time.Now(),
	}
}

// UpdateMetrics 메트릭 업데이트
func (pm *PrometheusMonitor) UpdateMetrics(ctx context.Context) {
	// 시스템 메트릭 업데이트
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	goroutinesGauge.Set(float64(runtime.NumGoroutine()))
	memoryAllocGauge.Set(float64(m.Alloc))
	memorySysGauge.Set(float64(m.Sys))
	
	// DB 메트릭 업데이트
	if sqlDB, err := pm.db.DB(); err == nil {
		stats := sqlDB.Stats()
		dbConnectionsGauge.Set(float64(stats.OpenConnections))
		dbConnectionsInUseGauge.Set(float64(stats.InUse))
		dbConnectionsIdleGauge.Set(float64(stats.Idle))
	}
	
	// Redis 메트릭 업데이트
	if length, err := pm.redisQueue.GetQueueLength(ctx); err == nil {
		redisQueueLengthGauge.Set(float64(length))
	}
}

// StartMonitoring 지속적인 모니터링 시작
func (pm *PrometheusMonitor) StartMonitoring(ctx context.Context) {
	ticker := time.NewTicker(15 * time.Second) // 프로메테우스 표준 간격
	defer ticker.Stop()
	
	log.Println("[PROMETHEUS] 메트릭 수집 시작")
	
	for {
		select {
		case <-ctx.Done():
			log.Println("[PROMETHEUS] 메트릭 수집 종료")
			return
		case <-ticker.C:
			pm.UpdateMetrics(ctx)
		}
	}
}

// RecordTransaction 트랜잭션 처리 시간 기록
func (pm *PrometheusMonitor) RecordTransaction(status string, duration time.Duration) {
	transactionProcessingTimeHistogram.Observe(duration.Seconds())
	transactionStatusCounter.WithLabelValues(status).Inc()
}

// GetUptime 업타임 반환
func (pm *PrometheusMonitor) GetUptime() time.Duration {
	return time.Since(pm.startTime)
} 