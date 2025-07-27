package database

import (
	"context"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
)

func WithTx[T any](ctx context.Context, dbQuery *query.Query, logger logger.Logger, fn func(context.Context, *query.QueryTx) (T, error)) (T, error) {
	return gorm.WithTx(ctx, dbQuery, logger, fn)
}
