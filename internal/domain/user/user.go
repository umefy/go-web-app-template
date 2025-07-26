package model

import (
	"time"

	dbModel "github.com/umefy/go-web-app-template/gorm/generated/model"
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	"github.com/umefy/go-web-app-template/pkg/null"
	"github.com/umefy/go-web-app-template/pkg/validation"
	"github.com/umefy/godash/sliceskit"
)

type User struct {
	ID        int
	Email     string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	Orders    []orderDomain.Order
}

func (u User) MapToDbModel() *dbModel.User {
	return &dbModel.User{
		Email:     null.ValueFrom(u.Email),
		Age:       null.ValueFrom(u.Age),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u User) CreateFromDbModel(dbUserModel *dbModel.User) *User {
	return &User{
		ID:        dbUserModel.ID,
		Email:     dbUserModel.Email.ValueOrZero(),
		Age:       dbUserModel.Age.ValueOrZero(),
		CreatedAt: dbUserModel.CreatedAt,
		UpdatedAt: dbUserModel.UpdatedAt,
		Orders: sliceskit.Map(dbUserModel.Orders, func(order dbModel.Order) orderDomain.Order {
			return *orderDomain.Order{}.CreateFromDbModel(&order)
		}),
	}
}

type UserCreateInput struct {
	Email string
	Age   int
}

func (u *UserCreateInput) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Age, validation.Max(100)),
	)
}

func (u UserCreateInput) MapToDbModel() *dbModel.User {
	return &dbModel.User{
		Email: null.ValueFrom(u.Email),
		Age:   null.ValueFrom(u.Age),
	}
}

type UserUpdateInput struct {
	Email *string
	Age   *int
}

func (u *UserUpdateInput) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email),
		validation.Field(&u.Age, validation.Max(100)),
	)
}

func (u UserUpdateInput) MapToDbModel() *dbModel.User {
	return &dbModel.User{
		Email: null.ValueFromPtr(u.Email),
		Age:   null.ValueFromPtr(u.Age),
	}
}
