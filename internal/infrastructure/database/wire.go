package database

import (
	"github.com/google/wire"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repository"
	gorm "github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm"
)

var WireSet = wire.NewSet(
	gorm.NewDB,
	gorm.NewDBQuery,
	gorm.NewUserRepository,
	wire.Bind(new(userRepo.Repository), new(*gorm.UserRepository)),
)
