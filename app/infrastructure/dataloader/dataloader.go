package dataloader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	gopher_dataloader "github.com/graph-gophers/dataloader"

	"github.com/o-t-k-t/graphl_app_trial/app/adapter/controller"
	"github.com/o-t-k-t/graphl_app_trial/ent"
	"github.com/o-t-k-t/graphl_app_trial/graph/model"
)

// functions for Resolver
// For returns the dataloader for a given context
func For(ctx context.Context) *DataLoader {
	return ctx.Value(loadersKey).(*DataLoader)
}

// DataLoader offers data loaders scoped to a context
type DataLoader struct {
	UserCarsLoader *dataloader.Loader
}

// newDataLoader returns the instantiated Loaders struct for use in a request
func newDataLoader(e *ent.Client) *DataLoader {
	return &DataLoader{
		UserCarsLoader: withBatchedLoader(controller.NewUserController(e).BatchFindCars),
	}
}

// LoadUserCars wraps the Cars dataloader for efficient retrieval by user ID
func (i *DataLoader) LoadUserCars(ctx context.Context, userID string) (*model.Cars, error) {
	thunk := i.UserCarsLoader.Load(ctx, gopher_dataloader.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.Cars), nil
}
