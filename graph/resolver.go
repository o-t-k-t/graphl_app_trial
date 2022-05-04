package graph

//go:generate go run github.com/99designs/gqlgen@v0.17.5 generate

import (
	"github.com/o-t-k-t/graphl_app_trial/ent"

	"github.com/o-t-k-t/graphl_app_trial/app/adapter/controller"
)

type Resolver struct {
	entClient ent.Client

	UserController controller.UserController
}
