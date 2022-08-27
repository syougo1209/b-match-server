.PHONY: down
down: ## down
	docker-compose down
build: ##build
	docker-compose build --no-cache

migrate: ## Migrate develop database
	mysqldef -u root -p '' -h 127.0.0.1 -P 3306 root < ./_tools/mysql/schema.sql

.PHONY: test
test: ## Execute tests
	go test -race -shuffle=on ./...
