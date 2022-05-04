package dataloader

import (
	"context"
	"net/http"

	"github.com/o-t-k-t/graphl_app_trial/ent"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// Middleware injects a DataLoader into the request context so it can be
// used later in the schema resolvers
func Middleware(entClient *ent.Client, next http.Handler) http.Handler {
	loader := newDataLoader(entClient)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loader)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}
