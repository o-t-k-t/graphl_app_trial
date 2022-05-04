generate:
	go generate ./...

build: generate
	go build -o ./bin cmd/server/server.go
	go build -o ./bin cmd/migration/migration.go

install:
	go intall github.com/99designs/gqlgen@v0.17.5

server:
	go run cmd/server/server.go

migration:
	go run cmd/migration/migration.go

schema:
	go run entgo.io/ent/cmd/ent describe ./ent/schema
