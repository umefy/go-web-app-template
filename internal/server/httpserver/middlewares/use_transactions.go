package middlewares

import (
	"context"
	"net/http"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	"github.com/umefy/go-web-app-template/internal/server/httpserver/handler"
)

type TransactionKey string

const (
	Transaction TransactionKey = "tx"
)

func UseTransaction(dbQuery *query.Query) handler.Middleware {
	return func(next handler.HandlerFunc) handler.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) error {
			tx := dbQuery.Begin()
			// nolint: errcheck
			defer tx.Rollback()

			ctx := context.WithValue(r.Context(), Transaction, tx)
			err := next(w, r.WithContext(ctx))
			if err != nil {
				return err
			}

			err = tx.Commit()
			if err != nil {
				return err
			}
			return nil
		}
	}
}
