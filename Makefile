.PHONY: all contracts-test backend-test docker-build docker-up docker-down contracts-node contracts-deploy

all: contracts-test backend-test

node_modules:
	[ -d node_modules ] || npm install

contracts-test: node_modules
	npx hardhat test

contracts-node: node_modules
	npx hardhat node
	@echo ""
	@echo "────────────────────────────────────────────"
	@echo "Hardhat 노드가 실행되었습니다."
	@echo "아래 첫 번째 Account의 Private Key를 복사해 .env의 MONSTER_GAME_PRIVKEY에 0x를 제거하고 붙여넣으세요."
	@echo "예시: MONSTER_GAME_PRIVKEY=abcdef123456... (0x 없이!)"
	@echo "────────────────────────────────────────────"
	@echo ""

contracts-deploy: node_modules
	@if ! lsof -i:8545 >/dev/null 2>&1; then \
		echo "Error: Hardhat node(localhost:8545)가 실행 중이어야 합니다. 먼저 'make contracts-node'로 노드를 실행하세요."; \
		exit 1; \
	fi
	npx hardhat run scripts/deploy.ts --network localhost
	@echo ""
	@echo "────────────────────────────────────────────"
	@echo "배포가 완료되었습니다."
	@echo "콘솔에 출력된 Proxy 주소를 .env의 MONSTER_GAME_CONTRACT에 복사해 넣으세요."
	@echo "예시: MONSTER_GAME_CONTRACT=0x1234abcd..."
	@echo "────────────────────────────────────────────"
	@echo ""

backend-test:
	go test ./backend/...

docker-build:
	docker-compose build

docker-up:
	docker-compose up --build -d
	@echo ""
	@echo "────────────────────────────────────────────"
	@echo "Swagger:  http://localhost:8080/swagger/index.html#/"
	@echo "Grafana:  http://localhost:3000/"
	@echo "────────────────────────────────────────────"
	@echo ""

docker-down:
	docker-compose down 