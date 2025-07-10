# 🧟‍♂️ Monster Hunt GameFi

블록체인 기반 GameFi 프로젝트로, 플레이어가 몬스터를 사냥하면 토큰 보상을 받는 시스템입니다.  
스마트 컨트랙트와 Golang 백엔드, Redis 큐/워커, Docker 기반 운영 자동화까지 **Play to Earn**의 실전 구조를 구현했습니다.

---

## 🚀 프로젝트 목표

- **GameFi 메커니즘 구현**: 몬스터 사냥 → 토큰 보상 → 잔고 조회 흐름 완성
- **Solidity 스마트컨트랙트**: ERC-20 기반 보상 토큰 및 몬스터 사냥 로직 작성
- **Golang 백엔드 서버**: 유저 등록, 사냥 요청, 보상 조회 API 제공 (실제 컨트랙트 연동)
- **Web3 연동**: 백엔드 ↔ 스마트컨트랙트 상호작용 완전 일원화
- **운영 자동화/확장성**: Redis 큐+워커, Docker/compose 기반 운영
- **보안 강화**: 그레이스풀 종료, 입력 검증, 동시성 제어, 에러 처리

---

## 🛠 기술 스택

| 분야 | 기술 |
|------|------|
| 스마트컨트랙트 | Solidity, Hardhat (Proxy/Upgradeable 패턴) |
| 백엔드 | Golang, go-ethereum, Gin, gorm |
| 큐/비동기 | Redis, worker(Go) |
| 트랜잭션 상태관리 | Postgres, gorm, confirmation worker |
| 배포/운영 | Docker, docker-compose, .env |
| 보안/안정성 | Rate Limiting, 그레이스풀 종료, 입력 검증, 동시성 제어 |
| 모니터링 | Prometheus, Grafana |
| 문서화 | Swagger (자동 생성) |

---

## 📦 디렉토리 구조 (최신)

```
monster-hunt-gamefi/
├── contracts/                # Solidity 스마트컨트랙트
│   ├── MonsterGame.sol      # 프록시(업그레이더블) + 보안(Pausable, ReentrancyGuard 등) 적용
│   ├── MonsterGameV2.sol    # 업그레이드 예시 (version 함수 추가)
│   └── MyGameToken.sol
├── scripts/                  # 배포 스크립트
│   └── deploy.ts            # import/export 스타일, 프록시 배포
├── test/                    # 테스트 코드
│   └── MonsterGame.test.ts  # 프록시 업그레이드/상태 유지/보안 테스트 포함
├── artifacts/                # 컴파일 산출물(ABI 등)
├── backend/
│   ├── application/          # 서비스 계층
│   ├── domain/               # 도메인 모델
│   ├── infrastructure/       # 이더리움/Redis/DB 연동 + 보안 유틸리티
│   │   ├── security.go      # 보안 유틸리티 (GetClientIP, SafeExecute, RetryWithBackoff 등)
│   │   ├── redis_queue.go   # Redis 큐 관리 (Close 메서드 추가)
│   │   └── ...
│   └── interface/            # API 라우터 및 Swagger
├── cmd/
│   ├── api/                  # API 서버 엔트리포인트 (그레이스풀 종료 지원)
│   ├── worker/               # 트랜잭션 제출 워커 (그레이스풀 종료 + 재시도 로직)
│   └── confirmation_worker/  # 트랜잭션 상태 확인 워커 (동시성 제어)
├── monitoring/               # Prometheus + Grafana 모니터링
│   ├── prometheus/
│   └── grafana/
├── Dockerfile
├── docker-compose.yml
├── .env                      # 환경변수 파일 (예시 포함)
└── README.md
```

---

## 🗄️ 데이터베이스 마이그레이션 구조

- **001_create_tx_status.sql**
  - `tx_status` 테이블 및 기본 인덱스 생성
  - 파티셔닝 없이 단일 테이블로 시작
- **002_partition_and_optimize_tx_status.sql**
  - `tx_status` 테이블을 파티셔닝 테이블로 변경 (PRIMARY KEY, UNIQUE 제약에 파티션 키 포함)
  - 월별 파티션 생성, 인덱스, 통계 뷰, 파티션 관리 함수 등 포함

### 📝 마이그레이션 적용 방법 예시

**Docker Compose 환경에서 Postgres 컨테이너에 접속 후:**
```bash
# Postgres 컨테이너 내부에서 실행
psql -U postgres -d monster_gamefi -f /docker-entrypoint-initdb.d/001_create_tx_status.sql
psql -U postgres -d monster_gamefi -f /docker-entrypoint-initdb.d/002_partition_and_optimize_tx_status.sql
```

**로컬 환경에서 직접 실행:**
```bash
psql -h localhost -U postgres -d monster_gamefi -f backend/infrastructure/migrations/001_create_tx_status.sql
psql -h localhost -U postgres -d monster_gamefi -f backend/infrastructure/migrations/002_partition_and_optimize_tx_status.sql
```

---

## ⚔️ 핵심 기능 및 구조

### 🎮 게임 로직 (온체인 연동)
- `registerPlayer(name)` : 플레이어 등록 (실제 컨트랙트 트랜잭션)
- `huntMonster(monsterId)` : 몬스터 사냥 → 보상 지급 (실제 컨트랙트 트랜잭션)
- `addMonster(name, hp, reward)` : 몬스터 추가 (owner만 가능)

### 🪙 토큰 시스템
- `MyGameToken` : ERC-20 기반 보상 토큰 발행
- 사냥 시 자동 토큰 전송

### 🏗️ 구조적 특징
- **MonsterGame 컨트랙트는 프록시(업그레이더블) 패턴 + 보안(Pausable, ReentrancyGuard 등)으로 배포**
- **pause/unpause(긴급 중지), 재진입 방지, 입력값 검증 등 보안 기능 내장**
- **상태/로직 분리, 안전한 업그레이드 지원**
- **API 서버 ↔ Redis 큐 ↔ worker ↔ 이더리움** 구조로 확장성/운영 자동화 실현
- **트랜잭션 상태 추적**: Postgres에 트랜잭션 상태(pending/success/fail) 저장, confirmation_worker가 receipt 확인 및 상태 업데이트
- **Swagger 문서 자동 생성/노출**: http://localhost:8080/swagger/index.html

### 🔒 보안 및 안정성 기능

#### **그레이스풀 종료**
- **시그널 처리**: SIGTERM/SIGINT 수신 시 안전한 종료
- **작업 완료 대기**: 진행 중인 트랜잭션 처리 완료 후 종료
- **리소스 정리**: DB/Redis 연결 안전한 종료

#### **입력 검증 강화**
- **Binding 태그**: `binding:"required,min=1,max=32"` 등 구조적 검증
- **추가 검증**: 비즈니스 로직 레벨 검증 (HP 범위, 보상 범위 등)
- **타입 안전성**: Go의 강타입 시스템 활용

#### **에러 처리 및 재시도**
- **SafeExecute**: 패닉 복구로 안정성 확보
- **RetryWithBackoff**: 지수 백오프로 일시적 실패 복구
- **컨텍스트 타임아웃**: 모든 외부 호출에 타임아웃 적용

#### **동시성 제어**
- **Redis BLPop**: 원자적 메시지 제거로 중복 처리 방지
- **조건부 업데이트**: Confirmation Worker에서 `WHERE status = 'pending'`으로 동시성 제어
- **고루틴 안전성**: WaitGroup으로 고루틴 관리

#### **보안 미들웨어**
- **GetClientIP**: 프록시 환경 고려한 클라이언트 IP 추출
- **IP 로깅**: 모든 요청에 대한 IP 로깅
- **프라이빗 IP 검증**: 선택적 접근 제한 가능

#### **Rate Limiting**
- **API 레벨**: Gin 미들웨어로 1분 1000회 제한
- **환경변수 제어**: `RATE_LIMIT_ENABLED`, `RATE_LIMIT_RATE`로 설정 가능

#### **헬스체크**
- **`/health` 엔드포인트**: 서비스 상태 확인
- **연결 상태 모니터링**: Redis/DB 연결 상태 확인

#### **모니터링**
- **Prometheus**: 메트릭 수집 및 노출
- **Grafana**: 대시보드 및 알림 설정
- **메트릭**: API 요청 수, 응답 시간, 에러율, 트랜잭션 상태 등

---

## ✅ 실행 및 개발/운영 방법

### 1. 스마트컨트랙트 컴파일 및 배포 (Hardhat)
```bash
npx hardhat compile
npx hardhat node # 새 터미널에서 실행, 종료하지 말 것
npx hardhat run scripts/deploy.ts --network localhost
```
- **MonsterGame은 프록시(업그레이더블) 패턴으로 배포**
- 배포 후 콘솔에 MonsterGame(프록시) 컨트랙트 주소와 배포 계정 프라이빗키가 출력됨

### 2. 프록시 업그레이드 및 테스트
- `contracts/MonsterGameV2.sol` : 업그레이드 예시 컨트랙트 (version 함수 추가)
- `test/MonsterGame.test.ts` :
  - 프록시 업그레이드(`upgrades.upgradeProxy`) 후 상태(state) 유지 및 신규 기능(version) 정상 동작 테스트
  - **pause/unpause, 입력값 검증, 권한, 보안 관련 테스트 포함**

### 3. 환경변수 설정 (.env 예시)
```env
MONSTER_GAME_RPC=http://localhost:8545
MONSTER_GAME_CONTRACT=0x...   # 프록시 주소로 설정
MONSTER_GAME_PRIVKEY=...
REDIS_URL=redis://localhost:6379
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=monster_gamefi
# Rate Limiting 설정
RATE_LIMIT_ENABLED=true
RATE_LIMIT_RATE=1000-M  # 1분 1000회
```

### 4. 전체 서비스 실행 (Docker Compose)
```bash
docker-compose up --build
```
- api, worker, confirmation_worker, redis, postgres, prometheus, grafana 컨테이너가 자동 실행
- .env 파일의 환경변수가 각 컨테이너에 주입됨
- Swagger: http://localhost:8080/swagger/index.html
- 헬스체크: http://localhost:8080/health
- Grafana: http://localhost:3000 (admin/admin)
- Prometheus: http://localhost:9090

### 5. 트랜잭션 상태 추적 구조
- API/worker가 트랜잭션 요청을 Redis 큐에 push
- worker가 트랜잭션 제출 후 Postgres에 상태(pending) 저장
- confirmation_worker가 pending 트랜잭션을 receipt로 확인, success/fail로 상태 업데이트
- 동시성 제어로 중복 처리 방지

### 6. 모니터링 대시보드
- **Grafana**: http://localhost:3000
  - API 메트릭: 요청 수, 응답 시간, 에러율
  - 트랜잭션 메트릭: 성공/실패율, 처리 시간
  - 시스템 메트릭: CPU, 메모리, 네트워크

### 7. API 테스트 예시
```bash
# 플레이어 등록
curl -X POST http://localhost:8080/players -H "Content-Type: application/json" -d '{"name":"Alice"}'

# 몬스터 추가 (owner만 가능)
curl -X POST http://localhost:8080/monsters -H "Content-Type: application/json" -d '{"name":"Goblin","hp":10,"reward":100}'

# 몬스터 사냥
curl -X POST http://localhost:8080/hunt -H "Content-Type: application/json" -d '{"monster_id":0}'

# 헬스체크
curl http://localhost:8080/health
```