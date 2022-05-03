package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"

	"github.com/o-t-k-t/graphl_app_trial/app/adapter/controller"
	"github.com/o-t-k-t/graphl_app_trial/app/adapter/repository"
	"github.com/o-t-k-t/graphl_app_trial/app/usecase"
	"github.com/o-t-k-t/graphl_app_trial/ent"
	"github.com/o-t-k-t/graphl_app_trial/graph"
	"github.com/o-t-k-t/graphl_app_trial/graph/generated"
)

const defaultPort = "8080"

func SetupServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Establish database connection.
	entClient, err := ent.Open("postgres", "dbname=feeder_development user=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer entClient.Close()

	// Setup GraphQL server.
	resolver := graph.Resolver{
		UserController: controller.UserController{
			usecase.UserUsecase{
				UserRepository: repository.UserRepository{
					EntClient: entClient,
				},
			},
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
