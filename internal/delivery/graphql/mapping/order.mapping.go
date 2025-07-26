package mapping

import (
	"strconv"
	"time"

	"github.com/umefy/go-web-app-template/internal/delivery/graphql/model"
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
)

func OrderModelToGraphqlOrder(order orderDomain.Order) *model.Order {
	return &model.Order{
		ID:        strconv.Itoa(order.ID),
		UserID:    strconv.Itoa(order.UserID),
		Amount:    order.Amount,
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
	}
}
