package model

import (
	"time"

	dbModel "github.com/umefy/go-web-app-template/gorm/generated/model"
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	"github.com/umefy/go-web-app-template/pkg/null"
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
