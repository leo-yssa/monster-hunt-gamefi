package infrastructure

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// RateLimiterConfig rate limiting 설정
type RateLimiterConfig struct {
	Enabled     bool
	RequestsPer int
	Window      time.Duration
	Burst       int
}

// GetClientIP 클라이언트 IP 주소 추출 (프록시 환경 고려)
func GetClientIP(r *http.Request) string {
	// X-Forwarded-For 헤더 확인 (Nginx 등 프록시)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// X-Real-IP 헤더 확인
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// 직접 연결인 경우
	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return ip
	}

	return r.RemoteAddr
}

// GenerateSecureToken 안전한 랜덤 토큰 생성
func GenerateSecureToken(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// ValidateIPAddress IP 주소 유효성 검사
func ValidateIPAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	}
	return true
}

// IsPrivateIP 프라이빗 IP 주소인지 확인
func IsPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	// 프라이빗 IP 범위 확인
	privateRanges := []struct {
		start, end net.IP
	}{
		{net.ParseIP("10.0.0.0"), net.ParseIP("10.255.255.255")},
		{net.ParseIP("172.16.0.0"), net.ParseIP("172.31.255.255")},
		{net.ParseIP("192.168.0.0"), net.ParseIP("192.168.255.255")},
		{net.ParseIP("127.0.0.0"), net.ParseIP("127.255.255.255")},
	}

	for _, r := range privateRanges {
		if inRange(parsedIP, r.start, r.end) {
			return true
		}
	}
	return false
}

func inRange(ip, start, end net.IP) bool {
	return bytes2Int(ip) >= bytes2Int(start) && bytes2Int(ip) <= bytes2Int(end)
}

func bytes2Int(ip net.IP) uint32 {
	ip = ip.To4()
	return uint32(ip[0])<<24 + uint32(ip[1])<<16 + uint32(ip[2])<<8 + uint32(ip[3])
}

// ConnectionPoolConfig 연결 풀 설정
type ConnectionPoolConfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

// DefaultConnectionPoolConfig 기본 연결 풀 설정
func DefaultConnectionPoolConfig() ConnectionPoolConfig {
	return ConnectionPoolConfig{
		MaxOpenConns:    100,
		MaxIdleConns:    25,
		ConnMaxLifetime: 5 * time.Minute,
		ConnMaxIdleTime: 5 * time.Minute,
	}
}

// GracefulShutdownConfig 그레이스풀 종료 설정
type GracefulShutdownConfig struct {
	ShutdownTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
}

// DefaultGracefulShutdownConfig 기본 그레이스풀 종료 설정
func DefaultGracefulShutdownConfig() GracefulShutdownConfig {
	return GracefulShutdownConfig{
		ShutdownTimeout: 30 * time.Second,
		ReadTimeout:     30 * time.Second,
		WriteTimeout:    30 * time.Second,
		IdleTimeout:     60 * time.Second,
	}
}

// CircuitBreakerConfig 서킷브레이커 설정
type CircuitBreakerConfig struct {
	FailureThreshold int
	RecoveryTimeout  time.Duration
	MaxRequests      int
}

// DefaultCircuitBreakerConfig 기본 서킷브레이커 설정
func DefaultCircuitBreakerConfig() CircuitBreakerConfig {
	return CircuitBreakerConfig{
		FailureThreshold: 5,
		RecoveryTimeout:  60 * time.Second,
		MaxRequests:      10,
	}
}

// ValidateContext 컨텍스트 유효성 검사
func ValidateContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("context cancelled: %v", ctx.Err())
	default:
		return nil
	}
}

// SafeExecute 안전한 함수 실행 (패닉 복구)
func SafeExecute(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()
	return fn()
}

// RetryWithBackoff 재시도 로직 (지수 백오프)
func RetryWithBackoff(ctx context.Context, fn func() error, maxRetries int, initialDelay time.Duration) error {
	var lastErr error
	delay := initialDelay

	for i := 0; i <= maxRetries; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := fn(); err == nil {
			return nil
		} else {
			lastErr = err
		}

		if i < maxRetries {
			time.Sleep(delay)
			delay *= 2 // 지수 백오프
		}
	}

	return fmt.Errorf("max retries exceeded: %v", lastErr)
} 