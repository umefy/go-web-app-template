package mapping

import (
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
)

func OrderModelToApiOrder(order *orderDomain.Order) api.Order {
	return api.Order{
		Id:        &order.ID,
		UserId:    order.UserID,
		Amount:    order.Amount,
		CreatedAt: &order.CreatedAt,
		UpdatedAt: &order.UpdatedAt,
	}
}
