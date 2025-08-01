package repo

import (
	"context"
	"errors"
	"log/slog"

	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	orderError "github.com/umefy/go-web-app-template/internal/domain/order/error"
	orderRepo "github.com/umefy/go-web-app-template/internal/domain/order/repo"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/query"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/repo/mapping"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	"github.com/umefy/godash/sliceskit"
	"gorm.io/gorm"
)

type OrderRepo struct {
	Logger  logger.Logger
	dbQuery *query.Query
}

var _ orderRepo.Repository = (*OrderRepo)(nil)

func NewOrderRepository(dbQuery *query.Query, logger logger.Logger) *OrderRepo {
	return &OrderRepo{Logger: logger, dbQuery: dbQuery}
}

func (r *OrderRepo) FindOrdersByUserID(ctx context.Context, userID int) ([]*orderDomain.Order, error) {
	orderQuery := r.dbQuery.Order
	orders, err := orderQuery.WithContext(ctx).Where(orderQuery.UserID.Eq(userID)).Find()
	if err != nil {
		r.Logger.ErrorContext(ctx, "FindOrdersByUserId error", slog.String("error", err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, orderError.OrderNotFound
		}
		return nil, err
	}
	return sliceskit.Map(orders, mapping.DbModelToDomainOrder), nil
}

func (r *OrderRepo) FindOrdersByUserIDs(ctx context.Context, userIDs []int) ([]*orderDomain.Order, error) {
	orderQuery := r.dbQuery.Order

	orders, err := orderQuery.WithContext(ctx).Where(orderQuery.UserID.In(userIDs...)).Find()
	if err != nil {
		r.Logger.ErrorContext(ctx, "FindOrdersByUserIDs error", slog.String("error", err.Error()))
		return nil, err
	}

	return sliceskit.Map(orders, mapping.DbModelToDomainOrder), nil
}
