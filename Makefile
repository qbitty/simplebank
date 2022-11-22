DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
postgres:
	docker run --name postgres14 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql doc/db.dbml --postgres -o doc/schema.sql

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

docker_network:
	docker network create bank-network

docker_run_server:
	docker run --name simplebank --network bank-networke DB_SOUECE=postgresql://root:secret@postgres14:5432/simple_bank?sslmode=disable -p 80:8080 simplebank:latest

mock:
	mockgen -destination db/mock/store.go -package mockdb pro.qbitty/simplebank/db/sqlc Store

proto:
	rm -f doc/swagger/*.swagger.json
	rm -f pb/*.go
	rm -f doc/statik/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc/

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock proto