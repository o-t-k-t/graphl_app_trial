package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"

	"github.com/o-t-k-t/graphl_app_trial/app/adapter/controller"
	"github.com/o-t-k-t/graphl_app_trial/app/infrastructure/dataloader"
	"github.com/o-t-k-t/graphl_app_trial/ent"
	"github.com/o-t-k-t/graphl_app_trial/graph"
	"github.com/o-t-k-t/graphl_app_trial/graph/generated"
)

const defaultPort = "8080"

var entClient *ent.Client

func newResolver(e *ent.Client) graph.Resolver {
	return graph.Resolver{
		UserController: controller.NewUserController(e),
	}
}

func SetupServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/query", createGraphQLHandler())

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	closeGraphQLHandler()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func createGraphQLHandler() http.Handler {
	// Establish database connection.
	var err error
	entClient, err = ent.Open("postgres", "dbname=feeder_development user=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Create GraphQL Reslover
	resolver := newResolver(entClient)

	// Setup GraphQL server.
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))
	return dataloader.Middleware(entClient, srv)
}

func closeGraphQLHandler() {
	entClient.Close()
}
