package mapping

import (
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	userSrv "github.com/umefy/go-web-app-template/internal/service/user"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
	"github.com/umefy/godash/sliceskit"
)

func UserModelToApiUser(user *userDomain.User) api.User {
	return api.User{
		Id:        &user.ID,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
		Orders: sliceskit.Map(user.Orders, func(order orderDomain.Order) api.Order {
			return OrderModelToApiOrder(&order)
		}),
	}
}

func ApiUserCreateToUserModelCreate(input *api.UserCreate) *userSrv.UserCreateInput {
	return &userSrv.UserCreateInput{
		Email: input.GetEmail(),
		Age:   input.GetAge(),
	}
}

func ApiUserUpdateToUserModelUpdate(input *api.UserUpdate) *userSrv.UserUpdateInput {
	return &userSrv.UserUpdateInput{
		Email: input.Email.Get(),
		Age:   input.Age.Get(),
	}
}
