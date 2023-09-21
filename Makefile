ifneq (,$(wildcard ./.env))
    include .env
    export
endif

migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down
sqlc:
	sqlc generate && echo "sqlc generated successfully"
test:
	go clean -testcache
	go test -v -cover ./...
server:
	rm -f ./bin/simplebank
	go build -o bin/simplebank .
	./bin/simplebank
mock:
	mockery -d
setupdeps:
	(command -v sqlc >/dev/null 2>&1 || go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.21.0) && echo "sqlc installed successfully"
	(command -v migrate >/dev/null 2>&1 || go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest) && echo "migrate installed successfully"
	(command -v mockery >/dev/null 2>&1 || go install github.com/vektra/mockery/v2@v2.33.3) && echo "mockery installed successfully"

.SILENT:
.PHONY: migrateup migratedown sqlc test server mock setupdeps
.DEFAULT_GOAL := server