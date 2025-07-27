package database

import (
	"github.com/google/wire"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repo"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/repo"
)

var WireSet = wire.NewSet(
	gorm.NewDB,
	gorm.NewDBQuery,
	repo.NewUserRepository,
	wire.Bind(new(userRepo.Repository), new(*repo.UserRepo)),
)
