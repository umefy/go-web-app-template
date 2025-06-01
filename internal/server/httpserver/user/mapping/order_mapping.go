package mapping

import (
	"time"

	"github.com/umefy/go-web-app-template/gorm/generated/model"
	api "github.com/umefy/go-web-app-template/openapi/protogen/v1/models"
)

func OrderModelToApiOrder(order *model.Order) *api.Order {
	return &api.Order{
		Id:        int32(order.ID),
		UserId:    int32(order.UserID.ValueOrZero()),
		Amount:    float32(order.Amount.ValueOrZero()),
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
	}
}
