package mapping

import (
	"strconv"
	"time"

	orderModel "github.com/umefy/go-web-app-template/internal/domain/order/model"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/graphql/model"
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
