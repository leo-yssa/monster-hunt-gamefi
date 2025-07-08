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

---

## 🛠 기술 스택

| 분야 | 기술 |
|------|------|
| 스마트컨트랙트 | Solidity, Hardhat |
| 백엔드 | Golang, go-ethereum, Gin, gorm |
| 큐/비동기 | Redis, worker(Go) |
| 트랜잭션 상태관리 | Postgres, gorm, confirmation worker |
| 배포/운영 | Docker, docker-compose, .env |
| 문서화 | Swagger (자동 생성) |

---

## 📦 디렉토리 구조 (최신)

```
monster-hunt-gamefi/
├── contracts/                # Solidity 스마트컨트랙트
│   ├── MonsterGame.sol
│   └── MyGameToken.sol
├── scripts/                  # 배포 스크립트
│   └── deploy.ts
├── artifacts/                # 컴파일 산출물(ABI 등)
├── backend/
│   ├── application/          # 서비스 계층
│   ├── domain/               # 도메인 모델
│   ├── infrastructure/       # 이더리움/Redis/DB 연동
│   └── interface/            # API 라우터 및 Swagger
├── cmd/
│   ├── api/                  # API 서버 엔트리포인트
│   ├── worker/               # 트랜잭션 제출 워커
│   └── confirmation_worker/  # 트랜잭션 상태 확인 워커
├── Dockerfile
├── docker-compose.yml
├── .env                      # 환경변수 파일 (예시 포함)
└── README.md
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
- **인메모리 저장소 완전 제거**: 모든 상태/로직은 이더리움 컨트랙트에 기록
- **API 서버 ↔ Redis 큐 ↔ worker ↔ 이더리움** 구조로 확장성/운영 자동화 실현
- **트랜잭션 상태 추적**: Postgres에 트랜잭션 상태(pending/success/fail) 저장, confirmation_worker가 receipt 확인 및 상태 업데이트
- **Swagger 문서 자동 생성/노출**: http://localhost:8080/swagger/index.html

---

## ✅ 실행 및 개발/운영 방법

### 1. 스마트컨트랙트 컴파일 및 배포 (Hardhat)
```bash
npx hardhat compile
npx hardhat node # 새 터미널에서 실행, 종료하지 말 것
npx hardhat run scripts/deploy.ts --network localhost
```
- 배포 후 콘솔에 MonsterGame 컨트랙트 주소와 배포 계정 프라이빗키가 출력됨

### 2. 환경변수 설정 (.env 예시)
```env
MONSTER_GAME_RPC=http://localhost:8545
MONSTER_GAME_CONTRACT=0x...
MONSTER_GAME_PRIVKEY=...
REDIS_URL=redis://localhost:6379
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=monster_gamefi
# POSTGRES_DSN=host=postgres port=5432 user=postgres password=postgres dbname=monster_gamefi sslmode=disable TimeZone=UTC
```
- 프라이빗키는 0x 없이 입력
- Hardhat 노드 실행 시 출력되는 첫 번째 계정 사용 권장

### 3. 전체 서비스 실행 (Docker Compose)
```bash
docker-compose up --build
```
- api, worker, confirmation_worker, redis, postgres 컨테이너가 자동 실행
- .env 파일의 환경변수가 각 컨테이너에 주입됨
- Swagger: http://localhost:8080/swagger/index.html

### 4. 트랜잭션 상태 추적 구조
- API/worker가 트랜잭션 요청을 Redis 큐에 push
- worker가 트랜잭션 제출 후 Postgres에 상태(pending) 저장
- confirmation_worker가 pending 트랜잭션을 receipt로 확인, success/fail로 상태 업데이트
- (확장) API에서 트랜잭션 상태 조회 엔드포인트 구현 가능

### 5. API 테스트 예시
```bash
curl -X POST http://localhost:8080/players -H "Content-Type: application/json" -d '{"name":"Alice"}'
curl -X POST http://localhost:8080/monsters -H "Content-Type: application/json" -d '{"name":"Goblin","hp":10,"reward":100}'
curl -X POST http://localhost:8080/hunt -H "Content-Type: application/json" -d '{"monster_id":0}'
```
- 모든 트랜잭션 요청은 Redis 큐에 push되고, worker가 실제 이더리움 트랜잭션을 처리하며, 상태는 Postgres에서 추적됨

---
