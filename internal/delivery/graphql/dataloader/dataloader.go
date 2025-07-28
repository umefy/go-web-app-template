package dataloader

import (
	"context"
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/graphql/model"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	"github.com/umefy/go-web-app-template/internal/service/order"
	"github.com/vikstrous/dataloadgen"
)

type ctxKey struct{}

var LoaderCtxKey ctxKey

type Loaders struct {
	OrderLoader *dataloadgen.Loader[string, []*model.Order] // userID -> orders
}

type LoaderDeps struct {
	OrderService order.Service
	Logger       logger.Logger
}

func NewLoaders(ctx context.Context, deps LoaderDeps) *Loaders {
	return &Loaders{
		OrderLoader: createOrderLoader(ctx, deps.Logger, deps.OrderService),
	}
}

// Middleware injects data loaders into the context
func Middleware(next http.Handler, deps LoaderDeps) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders(r.Context(), deps)
		r = r.WithContext(context.WithValue(r.Context(), LoaderCtxKey, loader))
		next.ServeHTTP(w, r)
	})
}
