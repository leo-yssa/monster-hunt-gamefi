package application

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulShutdown 애플리케이션 생명주기 관리
// HTTP 서버를 그레이스풀하게 종료합니다.
func GracefulShutdown(srv *http.Server, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("[APPLICATION] 서버 종료 신호 수신...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("[APPLICATION] 서버 강제 종료: %v", err)
	} else {
		log.Println("[APPLICATION] 서버 그레이스풀 종료 완료")
	}
}

// GracefulShutdownWithDefaultTimeout 기본 타임아웃(30초)으로 그레이스풀 종료
func GracefulShutdownWithDefaultTimeout(srv *http.Server) {
	GracefulShutdown(srv, 30*time.Second)
} 