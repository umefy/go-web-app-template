package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
)

func Transaction(dbQuery *database.Query, logger logger.Logger) handler.Middleware {
	return func(next handler.HandlerFunc) handler.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) (err error) {
			_, err = database.WithTx(r.Context(), dbQuery, logger, func(ctx context.Context, tx *database.QueryTx) (any, error) {
				ctx = context.WithValue(ctx, database.TransactionCtxKey, tx)
				return nil, next(w, r.WithContext(ctx))
			})
			return err
		}
	}
}

func GetTransaction(ctx context.Context) (*database.QueryTx, error) {
	tx, ok := ctx.Value(database.TransactionCtxKey).(*database.QueryTx)
	if !ok {
		return nil, errors.New("transaction not found")
	}
	return tx, nil
}
