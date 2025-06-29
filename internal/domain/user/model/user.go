package model

import (
	"time"

	dbModel "github.com/umefy/go-web-app-template/gorm/generated/model"
	orderModel "github.com/umefy/go-web-app-template/internal/domain/order/model"
	"github.com/umefy/go-web-app-template/pkg/null"
	"github.com/umefy/go-web-app-template/pkg/validation"
	"github.com/umefy/godash/sliceskit"
)

type User struct {
	ID        int
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	Orders    []orderModel.Order
}

func (u User) MapToDbModel() *dbModel.User {
	return &dbModel.User{
		Name:      null.ValueFrom(u.Name),
		Age:       null.ValueFrom(u.Age),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u User) CreateFromDbModel(dbUserModel *dbModel.User) *User {
	return &User{
		ID:        dbUserModel.ID,
		Name:      dbUserModel.Name.ValueOrZero(),
		Age:       dbUserModel.Age.ValueOrZero(),
		CreatedAt: dbUserModel.CreatedAt,
		UpdatedAt: dbUserModel.UpdatedAt,
		Orders: sliceskit.Map(dbUserModel.Orders, func(order dbModel.Order) orderModel.Order {
			return *orderModel.Order{}.CreateFromDbModel(&order)
		}),
	}
}

type UserCreateInput struct {
	Name string
	Age  int
}

func (u *UserCreateInput) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Age, validation.Max(100)),
	)
}

func (u UserCreateInput) MapToDbModel() *dbModel.User {
	return &dbModel.User{
		Name: null.ValueFrom(u.Name),
		Age:  null.ValueFrom(u.Age),
	}
}

type UserUpdateInput struct {
	Name *string
	Age  *int
}

func (u *UserUpdateInput) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name),
		validation.Field(&u.Age, validation.Max(100)),
	)
}

func (u UserUpdateInput) MapToDbModel() *dbModel.User {
	return &dbModel.User{
		Name: null.ValueFromPtr(u.Name),
		Age:  null.ValueFromPtr(u.Age),
	}
}
