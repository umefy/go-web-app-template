package mapping

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/umefy/go-web-app-template/gorm/generated/model"
	api "github.com/umefy/go-web-app-template/openapi/protogen/v1/models"
)

func UserModelToApiUser(user *model.User) *api.User {
	return &api.User{
		Id:        int32(user.ID),
		Name:      user.Name.ValueOrZero(),
		Age:       int32(user.Age.ValueOrZero()),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

func ApiCreateUserInputToUserModel(input *api.CreateUserInput) *model.User {
	return &model.User{
		Name: null.StringFrom(input.Name),
		Age:  null.IntFrom(int64(input.Age.GetValue())),
	}
}
