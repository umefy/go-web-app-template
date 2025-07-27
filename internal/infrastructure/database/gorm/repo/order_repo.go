package repo

import (
	"context"
	"errors"
	"log/slog"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	orderError "github.com/umefy/go-web-app-template/internal/domain/order/error"
	orderRepo "github.com/umefy/go-web-app-template/internal/domain/order/repo"
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

func (r *OrderRepo) FindOrdersByUserId(ctx context.Context, userId int) ([]*orderDomain.Order, error) {
	orderQuery := r.dbQuery.Order
	orders, err := orderQuery.WithContext(ctx).Where(orderQuery.UserID.Eq(userId)).Find()
	if err != nil {
		r.Logger.ErrorContext(ctx, "FindOrdersByUserId error", slog.String("error", err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, orderError.OrderNotFound
		}
		return nil, err
	}
	return sliceskit.Map(orders, mapping.DbModelToDomainOrder), nil
}

func (r *OrderRepo) FindOrdersByUserIds(ctx context.Context, userIds []int) ([]*orderDomain.Order, error) {
	orderQuery := r.dbQuery.Order

	orders, err := orderQuery.WithContext(ctx).Where(orderQuery.UserID.In(userIds...)).Find()
	if err != nil {
		return nil, err
	}
	return sliceskit.Map(orders, mapping.DbModelToDomainOrder), nil
}
