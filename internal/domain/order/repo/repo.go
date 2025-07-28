package repo

import (
	"context"

	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
)

type Repository interface {
	FindOrdersByUserID(ctx context.Context, userID int) ([]*orderDomain.Order, error)
	FindOrdersByUserIDs(ctx context.Context, userIDs []int) ([]*orderDomain.Order, error)
}
