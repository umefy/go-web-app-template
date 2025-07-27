package user

import (
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type UserUpdateInput struct {
	Email *string
	Age   *int
}

func (u *UserUpdateInput) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.IsEmail),
		validation.Field(&u.Age, validation.Min(0), validation.Max(100)),
	)
}

func updateDomainUser(user *userDomain.User, updateInput *UserUpdateInput) *userDomain.User {
	if updateInput.Email != nil {
		user.Email = *updateInput.Email
	}
	if updateInput.Age != nil {
		user.Age = *updateInput.Age
	}
	return user
}
