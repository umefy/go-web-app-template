package mapping

import (
	"strconv"
	"time"

	"github.com/umefy/go-web-app-template/internal/delivery/graphql/model"
	orderModel "github.com/umefy/go-web-app-template/internal/domain/order/model"
)

func OrderModelToGraphqlOrder(order orderModel.Order) *model.Order {
	return &model.Order{
		ID:        strconv.Itoa(order.ID),
		UserID:    strconv.Itoa(order.UserID),
		Amount:    order.Amount,
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
	}
}
