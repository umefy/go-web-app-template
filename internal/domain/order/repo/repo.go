package repo

import (
	"context"

	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
)

type Repository interface {
	FindOrdersByUserId(ctx context.Context, userId int) ([]*orderDomain.Order, error)
	FindOrdersByUserIds(ctx context.Context, userIds []int) ([]*orderDomain.Order, error)
}
