package dataloader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/o-t-k-t/graphl_app_trial/app/adapter/controller"
)

// withBatchedLoader translates dataloader specific definitions into app (contoller) definitions
func withBatchedLoader(f func(ctx context.Context, keys controller.DataloaderKeys) []*dataloader.Result) *dataloader.Loader {
	return dataloader.NewBatchedLoader(
		func(ctx context.Context, k dataloader.Keys) []*dataloader.Result {
			return f(ctx, DataloaderKeys(k))
		},
	)
}
