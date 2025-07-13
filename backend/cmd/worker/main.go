package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/config"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/domain"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/infrastructure"
	"gorm.io/datatypes"
)

type TxRequest struct {
	Action string                 `json:"action"` // "register", "hunt", "addMonster" 등
	Params map[string]interface{} `json:"params"`
	User   string                 `json:"user"`
}

type Worker struct {
	rdb            *redis.Client
	repo           domain.TxStatusRepository
	monsterGame    domain.MonsterGamePort
	curveLPToken   domain.CurveLPTokenPort
	curveLPStaking domain.CurveLPStakingPort
	privKeyHex     string
	queueName      string
	ctx            context.Context
	cancel         context.CancelFunc
	wg             sync.WaitGroup
	shutdown       chan struct{}
}

func NewWorker(rdb *redis.Client, repo domain.TxStatusRepository, monsterGame domain.MonsterGamePort, curveLPToken domain.CurveLPTokenPort, curveLPStaking domain.CurveLPStakingPort, privKeyHex string) *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Worker{
		rdb:            rdb,
		repo:           repo,
		monsterGame:    monsterGame,
		curveLPToken:   curveLPToken,
		curveLPStaking: curveLPStaking,
		privKeyHex:     privKeyHex,
		queueName:      "tx_queue",
		ctx:            ctx,
		cancel:         cancel,
		shutdown:       make(chan struct{}),
	}
}

func (w *Worker) Start() {
	log.Println("[WORKER] Redis 큐에서 트랜잭션 요청을 처리합니다...")
	
	// 시그널 핸들링
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 워커 고루틴 시작
	w.wg.Add(1)
	go w.processQueue()

	// 시그널 대기
	<-sigChan
	log.Println("[WORKER] 종료 신호 수신...")

	// 그레이스풀 종료
	w.GracefulShutdown()
}

func (w *Worker) processQueue() {
	defer w.wg.Done()

	for {
		select {
		case <-w.ctx.Done():
			log.Println("[WORKER] 컨텍스트 취소됨, 워커 종료")
			return
		case <-w.shutdown:
			log.Println("[WORKER] 셧다운 신호 수신, 워커 종료")
			return
		default:
			// Redis에서 메시지 가져오기 (5초 타임아웃)
			res, err := w.rdb.BLPop(w.ctx, 5*time.Second, w.queueName).Result()
			if err == redis.Nil {
				continue // 큐에 데이터 없음
			} else if err != nil {
				log.Printf("[WORKER] Redis BLPop 에러: %v", err)
				time.Sleep(1 * time.Second) // 에러 시 잠시 대기
				continue
			}
			if len(res) < 2 {
				continue
			}

			// 메시지 파싱
			var req TxRequest
			if err := json.Unmarshal([]byte(res[1]), &req); err != nil {
				log.Printf("[WORKER] TxRequest 파싱 실패: %v", err)
				continue
			}

			// 트랜잭션 처리 (고루틴으로 비동기 처리)
			w.wg.Add(1)
			go w.processTransaction(req)
		}
	}
}

func (w *Worker) processTransaction(req TxRequest) {
	defer w.wg.Done()

	log.Printf("[WORKER] 트랜잭션 요청 처리: %+v", req)

	// SafeExecute로 안전한 트랜잭션 처리
	err := infrastructure.SafeExecute(func() error {
		// 트랜잭션 제출 및 DB 저장
		txHash := ""
		errMsg := ""
		status := "pending"
		var paramsJSON datatypes.JSON
		if b, err := json.Marshal(req.Params); err == nil {
			paramsJSON = datatypes.JSON(b)
		}

		// RetryWithBackoff로 트랜잭션 처리 안정성 확보
		var err error
		err = infrastructure.RetryWithBackoff(w.ctx, func() error {
			switch req.Action {
			case "register":
				name, _ := req.Params["name"].(string)
				txHash, err = w.monsterGame.RegisterPlayer(w.ctx, name)
				if err != nil {
					errMsg = err.Error()
					status = "fail"
					return err
				} else {
					status = "pending"
				}
			case "addMonster":
				name, _ := req.Params["name"].(string)
				hp, _ := toInt(req.Params["hp"])
				reward, _ := toInt(req.Params["reward"])
				txHash, err = w.monsterGame.AddMonster(w.ctx, name, hp, reward)
				if err != nil {
					errMsg = err.Error()
					status = "fail"
					return err
				}
				status = "pending"
			case "hunt":
				monsterID, _ := toInt64(req.Params["monster_id"])
				txHash, err = w.monsterGame.HuntMonster(w.ctx, monsterID)
				if err != nil {
					errMsg = err.Error()
					status = "fail"
					return err
				} else {
					status = "pending"
				}
			case "stakeCurveLP":
				amount, _ := toInt64(req.Params["amount"])
				privKeyHex, _ := req.Params["privKeyHex"].(string)
				txHash, err = w.curveLPStaking.Stake(w.ctx, privKeyHex, amount)
				if err != nil {
					errMsg = err.Error()
					status = "fail"
					return err
				}
				status = "pending"
			case "unstakeCurveLP":
				amount, _ := toInt64(req.Params["amount"])
				privKeyHex, _ := req.Params["privKeyHex"].(string)
				txHash, err = w.curveLPStaking.Unstake(w.ctx, privKeyHex, amount)
				if err != nil {
					errMsg = err.Error()
					status = "fail"
					return err
				}
				status = "pending"
			default:
				return fmt.Errorf("알 수 없는 action: %s", req.Action)
			}
			return nil
		}, 3, 1*time.Second) // 최대 3회 재시도, 1초 초기 지연

		if err != nil {
			log.Printf("[WORKER] 트랜잭션 처리 실패: %v", err)
			return err
		}

		// DB에 상태 저장 (컨텍스트 타임아웃 적용)
		dbCtx, dbCancel := context.WithTimeout(w.ctx, 10*time.Second)
		defer dbCancel()

		var submittedBlock uint64 = 0
		if txHash != "" {
			// 트랜잭션 전송 후 현재 블록 넘버 조회 (MonsterGameAdapter 타입일 때만)
			if mgAdapter, ok := w.monsterGame.(*infrastructure.MonsterGameAdapter); ok {
				blockNum, err := mgAdapter.Client().BlockNumber(w.ctx)
				if err == nil {
					submittedBlock = blockNum
				}
			}
			w.repo.CreateTxStatus(dbCtx, &domain.TxStatus{
				TxHash:  txHash,
				Action:  req.Action,
				Params:  string(paramsJSON),
				UserID:  req.User,
				Status:  status,
				SubmittedBlock: submittedBlock,
			})
		} else if errMsg != "" {
			w.repo.CreateTxStatus(dbCtx, &domain.TxStatus{
				TxHash:  "",
				Action:  req.Action,
				Params:  string(paramsJSON),
				UserID:  req.User,
				Status:  status,
				SubmittedBlock: 0,
			})
		}

		return nil
	})

	if err != nil {
		log.Printf("[WORKER] SafeExecute 실패: %v", err)
	}
}

func (w *Worker) GracefulShutdown() {
	log.Println("[WORKER] 그레이스풀 종료 시작...")

	// 컨텍스트 취소로 진행 중인 작업들 중단 신호
	w.cancel()

	// 셧다운 신호 전송
	close(w.shutdown)

	// 모든 고루틴 완료 대기 (30초 타임아웃)
	done := make(chan struct{})
	go func() {
		w.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("[WORKER] 모든 작업 완료")
	case <-time.After(30 * time.Second):
		log.Println("[WORKER] 타임아웃으로 강제 종료")
	}

	// 리소스 정리
	if w.rdb != nil {
		w.rdb.Close()
	}

	log.Println("[WORKER] 그레이스풀 종료 완료")
}

func main() {
	_ = godotenv.Load()

	rdb := config.InitRedisClientFromEnv()
	db, err := config.InitGormDBFromEnv()
	if err != nil {
		log.Fatalf("Postgres DB 초기화 실패: %v", err)
	}
	repo := infrastructure.NewTxStatusRepository(db)

	// 환경변수에서 주소 읽기
	rpcURL := config.LoadEnvOrDefault("MONSTER_GAME_RPC", "http://localhost:8545")
	contractAddr := os.Getenv("MONSTER_GAME_CONTRACT")
	privKeyHex := os.Getenv("MONSTER_GAME_PRIVKEY")
	curveLPTokenAddr := os.Getenv("CURVE_LP_TOKEN_CONTRACT")
	curveLPStakingAddr := os.Getenv("CURVE_LP_STAKING_CONTRACT")
	if contractAddr == "" || privKeyHex == "" || curveLPTokenAddr == "" || curveLPStakingAddr == "" {
		log.Fatal("컨트랙트 주소와 프라이빗키 환경변수(MONSTER_GAME_CONTRACT, MONSTER_GAME_PRIVKEY, CURVE_LP_TOKEN_CONTRACT, CURVE_LP_STAKING_CONTRACT)가 필요합니다.")
	}

	client, err := ethclient.DialContext(context.Background(), rpcURL)
	if err != nil {
		log.Fatalf("이더리움 클라이언트 연결 실패: %v", err)
	}

	monsterGameAdapter, _ := infrastructure.NewMonsterGameAdapter(
		client, contractAddr, privKeyHex,
	)
	curveLPTokenAdapter, _ := infrastructure.NewCurveLPTokenAdapter(
		client, curveLPTokenAddr,
	)
	curveLPStakingAdapter, _ := infrastructure.NewCurveLPStakingAdapter(
		client, curveLPStakingAddr,
	)

	worker := NewWorker(rdb, repo, monsterGameAdapter, curveLPTokenAdapter, curveLPStakingAdapter, privKeyHex)
	worker.Start()
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