ifneq (,$(wildcard ./.env))
    include .env
    export
else
	DB_SOURCE=postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable
endif

migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down
sqlc:
	sqlc generate
test:
	go clean -testcache
	go test -v -cover ./...
server:
	rm -f ./bin/simplebank
	go build -o bin/simplebank .
	./bin/simplebank
mockgen:
	mockgen -package mockdb -destination db/mock/store.go github.com/fauzanfebrian/simplebank/db/sqlc Store

.PHONY: migrateup migratedown sqlc test server mockgen