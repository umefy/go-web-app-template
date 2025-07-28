package mapping

import (
	api "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/generated"
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
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
