generate:
	go generate ./...

build:
	go build -o ./bin cmd/server/server.go
	go build -o ./bin cmd/migration/migration.go

server:
	go run cmd/server/server.go

migration:
	go run cmd/migration/migration.go

schema:
	go run entgo.io/ent/cmd/ent describe ./ent/schema
