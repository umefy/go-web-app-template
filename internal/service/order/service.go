package order

import (
	"context"
	"log/slog"

	domainOrder "github.com/umefy/go-web-app-template/internal/domain/order"
	"github.com/umefy/go-web-app-template/internal/domain/order/repo"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
)

type Service interface {
	GetOrdersByUserID(ctx context.Context, userID int) ([]*domainOrder.Order, error)
	GetOrdersByUserIDs(ctx context.Context, userIDs []int) ([]*domainOrder.Order, error)
}

type orderService struct {
	logger    logger.Logger
	orderRepo repo.Repository
}

var _ Service = (*orderService)(nil)

func NewService(logger logger.Logger, orderRepo repo.Repository) *orderService {
	return &orderService{
		logger:    logger,
		orderRepo: orderRepo,
	}
}

// GetOrdersByUserID implements Service.
func (s *orderService) GetOrdersByUserID(ctx context.Context, userID int) ([]*domainOrder.Order, error) {
	orders, err := s.orderRepo.FindOrdersByUserID(ctx, userID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get orders by user ID", slog.String("error", err.Error()))
		return nil, err
	}

	return orders, nil
}

// GetOrdersByUserIDs implements Service.
func (s *orderService) GetOrdersByUserIDs(ctx context.Context, userIDs []int) ([]*domainOrder.Order, error) {
	orders, err := s.orderRepo.FindOrdersByUserIDs(ctx, userIDs)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get orders by user IDs", slog.String("error", err.Error()))
		return nil, err
	}
	return orders, nil
}
