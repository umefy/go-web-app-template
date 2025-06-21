package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	loggerSrv "github.com/umefy/go-web-app-template/internal/domain/logger/service"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/handler"
)

type transactionKey struct{}

var TransactionCtxKey = transactionKey{}

func Transaction(dbQuery *query.Query, loggerService loggerSrv.Service) handler.Middleware {
	return func(next handler.HandlerFunc) handler.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) (err error) {
			tx := dbQuery.Begin()
			defer func() {

				if rec := recover(); rec != nil {
					loggerService.ErrorContext(r.Context(), "Transaction rollback because of panic")
					//nolint:errcheck
					tx.Rollback()
					panic(rec)
				}

				if err != nil {
					loggerService.ErrorContext(r.Context(), "Transaction rollback", slog.String("error", err.Error()))
					//nolint:errcheck
					tx.Rollback()
				}
			}()

			ctx := context.WithValue(r.Context(), TransactionCtxKey, tx)
			err = next(w, r.WithContext(ctx))
			if err != nil {
				return err
			}

			return tx.Commit()
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
