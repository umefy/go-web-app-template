package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
)

type transactionKey struct{}

var TransactionCtxKey = transactionKey{}

func Transaction(dbQuery *query.Query, logger logger.Logger) handler.Middleware {
	return func(next handler.HandlerFunc) handler.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) (err error) {
			_, err = database.WithTx(r.Context(), dbQuery, logger, func(ctx context.Context, tx *query.QueryTx) (any, error) {
				ctx = context.WithValue(ctx, TransactionCtxKey, tx)
				return nil, next(w, r.WithContext(ctx))
			})
			return err
		}
	}
}

func GetTransaction(ctx context.Context) (*query.QueryTx, error) {
	tx, ok := ctx.Value(TransactionCtxKey).(*query.QueryTx)
	if !ok {
		return nil, errors.New("transaction not found")
	}
	return tx, nil
}
