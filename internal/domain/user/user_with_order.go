package user

import (
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
)

type UserWithOrder struct {
	User
	Orders []orderDomain.Order
}
