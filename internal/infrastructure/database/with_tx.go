package database

import (
	"context"
	"log/slog"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	loggerSrv "github.com/umefy/go-web-app-template/internal/domain/logger/service"
)

func WithTx[T any](ctx context.Context, dbQuery *query.Query, loggerService loggerSrv.Service, fn func(context.Context, *query.QueryTx) (T, error)) (T, error) {
	tx := dbQuery.Begin()
	var err error
	defer func() {
		if rec := recover(); rec != nil {
			loggerService.ErrorContext(ctx, "Transaction rollback because of panic")
			//nolint:errcheck
			tx.Rollback()
			panic(rec)
		}

		if err != nil {
			loggerService.ErrorContext(ctx, "Transaction rollback", slog.String("error", err.Error()))
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
