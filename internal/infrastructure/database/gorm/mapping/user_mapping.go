package mapping

import (
	dbModel "github.com/umefy/go-web-app-template/gorm/generated/model"
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	"github.com/umefy/go-web-app-template/pkg/null"
	"github.com/umefy/godash/sliceskit"
)

func UserDbModelToUserDomain(user *dbModel.User) *userDomain.User {
	return &userDomain.User{
		ID:        user.ID,
		Email:     user.Email.ValueOrZero(),
		Age:       user.Age.ValueOrZero(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Orders: sliceskit.Map(user.Orders, func(order dbModel.Order) orderDomain.Order {
			return *OrderDbModelToOrderDomain(&order)
		}),
	}
}

func UserDomainToUserDbModel(user *userDomain.User) *dbModel.User {
	return &dbModel.User{
		ID:        user.ID,
		Email:     null.ValueFrom(user.Email),
		Age:       null.ValueFrom(user.Age),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
