package mapping

import (
	dbModel "github.com/umefy/go-web-app-template/gorm/generated/model"
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
)

func DbModelToDomainOrder(order *dbModel.Order) *orderDomain.Order {
	return &orderDomain.Order{
		ID:        order.ID,
		UserID:    order.UserID,
		Amount:    order.Amount.ValueOrZero(),
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}
