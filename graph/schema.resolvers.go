package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/o-t-k-t/graphl_app_trial/app/infrastructure/dataloader"
	"github.com/o-t-k-t/graphl_app_trial/graph/generated"
	"github.com/o-t-k-t/graphl_app_trial/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	return r.UserController.Create(ctx, input)
}

func (r *queryResolver) ListUsers(ctx context.Context) ([]*model.User, error) {
	return r.UserController.List(ctx)
}

func (r *userResolver) Cars(ctx context.Context, obj *model.User) (*model.Cars, error) {
	return dataloader.For(ctx).LoadUserCars(ctx, obj.ID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
