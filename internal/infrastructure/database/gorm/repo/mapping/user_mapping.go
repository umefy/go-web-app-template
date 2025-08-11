package mapping

import (
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	dbModel "github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/model"
	"github.com/umefy/go-web-app-template/pkg/null"
	"github.com/umefy/godash/sliceskit"
)

func DbModelToDomainUser(user *dbModel.User) *userDomain.User {
	return &userDomain.User{
		ID:        user.ID,
		Email:     user.Email.ValueOrZero(),
		Age:       user.Age.ValueOrZero(),
		Version:   user.Version,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func DomainUserToDbModel(user *userDomain.User) *dbModel.User {
	return &dbModel.User{
		ID:        user.ID,
		Email:     null.ValueFrom(user.Email),
		Age:       null.ValueFrom(user.Age),
		Version:   user.Version,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func DbModelToDomainUserWithOrder(user *dbModel.User, orders []dbModel.Order) *userDomain.UserWithOrder {
	return &userDomain.UserWithOrder{
		User: *DbModelToDomainUser(user),
		Orders: sliceskit.Map(orders, func(order dbModel.Order) orderDomain.Order {
			return *DbModelToDomainOrder(&order)
		}),
	}
}
