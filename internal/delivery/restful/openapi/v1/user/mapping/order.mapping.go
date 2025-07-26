package mapping

import (
	"github.com/umefy/go-web-app-template/internal/domain/order/model"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
)

func OrderModelToApiOrder(order *model.Order) api.Order {
	return api.Order{
		Id:        &order.ID,
		UserId:    order.UserID,
		Amount:    order.Amount,
		CreatedAt: &order.CreatedAt,
		UpdatedAt: &order.UpdatedAt,
	}
}
