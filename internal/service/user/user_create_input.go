package user

import (
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	"github.com/umefy/go-web-app-template/pkg/validation"
)

type UserCreateInput struct {
	Email string
	Age   int
}

func (u *UserCreateInput) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, validation.IsEmail),
		validation.Field(&u.Age, validation.Min(0), validation.Max(100)),
	)
}

func (u *UserCreateInput) MapToDomainUser() *userDomain.User {
	return &userDomain.User{
		Email: u.Email,
		Age:   u.Age,
	}
}
