package mapping

import (
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	userSrv "github.com/umefy/go-web-app-template/internal/service/user"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
)

func UserModelToApiUser(user *userDomain.User) api.User {
	return api.User{
		Id:        &user.ID,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
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
