package mapping

import (
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	dbModel "github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/model"
)

func DbModelToDomainOrder(order *dbModel.Order) *orderDomain.Order {
	return &orderDomain.Order{
		ID:          order.ID,
		UserID:      order.UserID,
		AmountCents: order.AmountCents.ValueOrZero(),
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
}
