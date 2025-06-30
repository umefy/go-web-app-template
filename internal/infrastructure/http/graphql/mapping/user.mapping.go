package mapping

import (
	"strconv"
	"time"

	userModel "github.com/umefy/go-web-app-template/internal/domain/user/model"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/graphql/model"
	"github.com/umefy/godash/sliceskit"
)

func UserModelToGraphqlUser(user *userModel.User) *model.User {
	return &model.User{
		ID:        strconv.Itoa(user.ID),
		Email:     user.Email,
		Age:       int32(user.Age),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		Orders:    sliceskit.Map(user.Orders, OrderModelToGraphqlOrder),
	}
}
