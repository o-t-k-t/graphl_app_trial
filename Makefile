# Regenerate auto-generated code.
generate:
	go generate ./...

# Geneerate unit test code.
# PATH=./...
# generate_ut:
# 	go install github.com/cweill/gotests/...
# 	gotests -all $(PATH)

# Build Binary for each entrypoints.
build: generate
	go build -o ./bin cmd/server/server.go
	go build -o ./bin cmd/migration/migration.go

# Install CLI tools.
install:
	go install github.com/99designs/gqlgen@v0.17.5
	go install entgo.io/ent/cmd/ent@v0.10.1

# Run local server.
server:
	go run cmd/server/server.go

# Run Database migration.
migration:
	go run cmd/migration/migration.go

# Show database sdchema.
schema:
	go run entgo.io/ent/cmd/ent@v0.10.1 describe ./ent/schema
