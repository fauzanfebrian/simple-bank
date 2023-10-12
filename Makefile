ifneq (,$(wildcard ./.env))
    include .env
    export
endif

migratecreate:
	migrate create -ext sql -dir db/migration -seq $(name)
migrateforce:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose force 1
migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down
migrateup1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1
migratedown1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

sqlc:
	sqlc generate && echo "sqlc generated successfully"

test:
	go clean -testcache
	GIN_MODE=test go test -v -cover ./...

server:
	rm -f ./bin/simplebank
	go build -o bin/simplebank .
	./bin/simplebank

mock:
	mockery -d

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out pb --grpc-gateway_opt paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.SILENT:
.PHONY: migrateup migratedown sqlc test server mock migratecreate migrateup1 migratedown1 migrateforce migratecreate db_docs db_schema proto evans
.DEFAULT_GOAL := server