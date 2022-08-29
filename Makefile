.PHONY: down
down: ## down
	docker-compose down --remove-orphans
build: ##build
	docker-compose build --no-cache

migrate: ## Migrate develop database
	mysqldef -u root -proot -h 127.0.0.1 -P 3306 b-match < ./_tools/mysql/schema.sql

.PHONY: test
test: ## Execute tests
	go test -race -shuffle=on ./...
