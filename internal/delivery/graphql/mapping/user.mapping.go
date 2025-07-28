package mapping

import (
	"strconv"
	"time"

	"github.com/umefy/go-web-app-template/internal/delivery/graphql/model"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
)

func DomainUserToGraphqlUser(user *userDomain.User) *model.User {
	return &model.User{
		ID:        strconv.Itoa(user.ID),
		Email:     user.Email,
		Age:       int32(user.Age),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}
