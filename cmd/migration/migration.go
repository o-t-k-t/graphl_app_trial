package main

import (
	"context"
	"log"

	"github.com/o-t-k-t/graphl_app_trial/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "dbname=feeder_development user=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// TODO: Versioned migration.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
