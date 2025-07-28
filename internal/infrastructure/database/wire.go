package database

import (
	"github.com/google/wire"
	orderRepo "github.com/umefy/go-web-app-template/internal/domain/order/repo"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repo"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/repo"
)

var WireSet = wire.NewSet(
	gorm.NewDB,
	gorm.NewDBQuery,
	repo.NewUserRepository,
	repo.NewOrderRepository,
	wire.Bind(new(userRepo.Repository), new(*repo.UserRepo)),
	wire.Bind(new(orderRepo.Repository), new(*repo.OrderRepo)),
)
