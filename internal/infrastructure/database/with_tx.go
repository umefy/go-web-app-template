package database

import (
	"context"
	"log/slog"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
)

func WithTx[T any](ctx context.Context, dbQuery *query.Query, logger logger.Logger, fn func(context.Context, *query.QueryTx) (T, error)) (T, error) {
	tx := dbQuery.Begin()
	var err error
	defer func() {
		if rec := recover(); rec != nil {
			logger.ErrorContext(ctx, "Transaction rollback because of panic")
			//nolint:errcheck
			tx.Rollback()
			panic(rec)
		}

		if err != nil {
			logger.ErrorContext(ctx, "Transaction rollback", slog.String("error", err.Error()))
			//nolint:errcheck
			tx.Rollback()
		}
	}()

	v, err := fn(ctx, tx)
	if err != nil {
		return v, err
	}

	return v, tx.Commit()
}
