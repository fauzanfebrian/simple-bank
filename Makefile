ifneq (,$(wildcard ./.env))
    include .env
    export
endif

migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	rm -f ./bin/simplebank
	go build -o bin/simplebank .
	./bin/simplebank

.PHONY: migrateup migratedown sqlc test server