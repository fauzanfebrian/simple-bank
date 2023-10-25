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
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--experimental_allow_proto3_optional \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out pb --grpc-gateway_opt paths=source_relative \
		--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simplebank \
		proto/*.proto
	statik -src=./doc/swagger -dest=./doc -f

evans:
	evans --host localhost --port 9090 -r repl

setup:
	echo "setup packages..."
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest || true;
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway || true;
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 || true;
	go install github.com/rakyll/statik || true;
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.21.0 || true;
	go install github.com/vektra/mockery/v2@v2.33.3 || true;
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc || true;
	go install google.golang.org/protobuf/cmd/protoc-gen-go || true;

.SILENT:
.PHONY: migrateup migratedown sqlc test server mock migratecreate migrateup1 migratedown1 migrateforce migratecreate db_docs db_schema proto evans setup
.DEFAULT_GOAL := server