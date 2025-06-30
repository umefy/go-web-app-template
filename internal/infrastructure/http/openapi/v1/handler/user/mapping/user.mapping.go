package mapping

import (
	orderModel "github.com/umefy/go-web-app-template/internal/domain/order/model"
	userModel "github.com/umefy/go-web-app-template/internal/domain/user/model"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
	"github.com/umefy/godash/sliceskit"
)

func UserModelToApiUser(user *userModel.User) api.User {
	return api.User{
		Id:        &user.ID,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
		Orders: sliceskit.Map(user.Orders, func(order orderModel.Order) api.Order {
			return OrderModelToApiOrder(&order)
		}),
	}
}

func ApiUserCreateToUserModelCreate(input *api.UserCreate) *userModel.UserCreateInput {
	return &userModel.UserCreateInput{
		Email: input.GetEmail(),
		Age:   input.GetAge(),
	}
}

func ApiUserUpdateToUserModelUpdate(input *api.UserUpdate) *userModel.UserUpdateInput {
	return &userModel.UserUpdateInput{
		Email: input.Email.Get(),
		Age:   input.Age.Get(),
	}
}
