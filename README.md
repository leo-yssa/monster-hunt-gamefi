# 🧟‍♂️ Monster Hunt GameFi

블록체인 기반 GameFi 프로젝트로, 플레이어가 몬스터를 사냥하면 토큰 보상을 받는 시스템입니다.  
스마트 컨트랙트와 Golang 백엔드를 활용해 **Play to Earn**의 핵심 개념을 구현했습니다.

---

## 🚀 프로젝트 목표

- **GameFi 메커니즘 구현**: 몬스터 사냥 → 토큰 보상 → 잔고 조회 흐름 완성
- **Solidity 스마트컨트랙트**: ERC-20 기반 보상 토큰 및 몬스터 사냥 로직 작성
- **Golang 백엔드 서버**: 유저 등록, 사냥 요청, 보상 조회 API 제공
- **Web3 연동**: 백엔드 ↔ 스마트컨트랙트 상호작용 구현

---

## 🛠 기술 스택

| 분야 | 기술 |
|------|------|
| 스마트컨트랙트 | Solidity, Hardhat |
| 백엔드 | Golang, go-ethereum, Gin |
| 블록체인 네트워크 | Local Hardhat Node, (확장 시 Testnet) |
| 배포 | Github 공개 코드 제출용 |

---

## 📦 디렉토리 구조 (예시)

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
│   ├── infrastructure/       # 이더리움 연동
│   ├── interface/            # API 라우터 및 Swagger
│   └── main.go               # 서버 실행 엔트리포인트
└── ...
```

---

## ⚔️ 핵심 기능

### 🎮 게임 로직
- `registerPlayer(name)` : 플레이어 등록 (Go/컨트랙트 모두 address는 트랜잭션 서명 계정 기준)
- `huntMonster(monsterId)` : 몬스터 사냥 → 보상 지급
- `getReward(address)` : (추후 구현) 보상 확인

### 🪙 토큰 시스템
- `MyGameToken` : ERC-20 기반 보상 토큰 발행
- 사냥 시 자동 토큰 전송

---

## ✅ 실행 및 개발 방법

### 1. 스마트컨트랙트 컴파일 및 배포 (로컬 Hardhat)
```bash
npx hardhat compile
npx hardhat node # 새 터미널에서 실행, 종료하지 말 것
npx hardhat run scripts/deploy.ts --network localhost
```
- 배포 후 콘솔에 MonsterGame 컨트랙트 주소와 배포 계정 프라이빗키가 출력됨

### 2. 환경변수 설정 (Go 연동용)
```bash
export MONSTER_GAME_RPC=http://localhost:8545
export MONSTER_GAME_CONTRACT=<MonsterGame_컨트랙트_주소>
export MONSTER_GAME_PRIVKEY=<Hardhat_테스트_계정_프라이빗키>
```
- 프라이빗키는 0x 없이 입력
- Hardhat 노드 실행 시 출력되는 첫 번째 계정 사용 권장

### 3. Swagger 문서 자동생성 및 서버 실행
```bash
export PATH=$PATH:$HOME/go/bin
swag init -g backend/main.go
# 의존성 설치 (최초 1회)
go get github.com/gin-gonic/gin github.com/swaggo/files github.com/swaggo/gin-swagger
# 서버 실행
go run ./backend/main.go
```
- Swagger UI: http://localhost:8080/swagger/index.html
- API 문서가 정상적으로 노출됨

### 4. API 테스트 예시
```bash
curl -X POST http://localhost:8080/players -H "Content-Type: application/json" -d '{"name":"Alice"}'
curl -X POST http://localhost:8080/monsters -H "Content-Type: application/json" -d '{"name":"Goblin","hp":10,"reward":100}'
curl -X POST http://localhost:8080/hunt -H "Content-Type: application/json" -d '{"monster_id":0}'
```

---
