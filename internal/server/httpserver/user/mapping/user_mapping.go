package mapping

import (
	"time"

	"github.com/umefy/go-web-app-template/gorm/generated/model"
	api "github.com/umefy/go-web-app-template/openapi/protogen/v1/models"
	"github.com/umefy/go-web-app-template/pkg/null"
	"github.com/umefy/godash/sliceskit"
)

func UserModelToApiUser(user *model.User) *api.User {
	return &api.User{
		Id:        int32(user.ID),
		Name:      user.Name.ValueOrZero(),
		Age:       int32(user.Age.ValueOrZero()),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		Orders: sliceskit.Map(user.Orders, func(order model.Order) *api.Order {
			return OrderModelToApiOrder(&order)
		}),
	}
}

func ApiCreateUserInputToUserModel(input *api.CreateUserInput) *model.User {
	return &model.User{
		Name: null.ValueFrom(input.Name),
		Age:  null.ValueFrom(int(input.Age)),
	}
}

func ApiUpdateUserInputToUserModel(input *api.UpdateUserInput) *model.User {
	return &model.User{
		Name: null.ValueFromWrapperspbString(input.Name),
		Age:  null.ValueFromWrapperspbInt32(input.Age),
	}
}
