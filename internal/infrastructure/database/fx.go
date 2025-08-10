package database

import (
	orderRepo "github.com/umefy/go-web-app-template/internal/domain/order/repo"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repo"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/repo"
	"go.uber.org/fx"
)

var Module = fx.Module("database",
	fx.Provide(
		gorm.NewDB,
		gorm.NewDBQuery,
		fx.Annotate(
			repo.NewUserRepository,
			fx.As(new(userRepo.Repository)),
		),
		fx.Annotate(
			repo.NewOrderRepository,
			fx.As(new(orderRepo.Repository)),
		),
	),
)
