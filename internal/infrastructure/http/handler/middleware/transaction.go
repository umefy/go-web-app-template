package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	loggerSrv "github.com/umefy/go-web-app-template/internal/domain/logger/service"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/handler"
)

type transactionKey struct{}

var TransactionCtxKey = transactionKey{}

func Transaction(dbQuery *query.Query, loggerService loggerSrv.Service) handler.Middleware {
	return func(next handler.HandlerFunc) handler.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) (err error) {
			_, err = database.WithTx(r.Context(), dbQuery, loggerService, func(ctx context.Context, tx *query.QueryTx) (any, error) {
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
