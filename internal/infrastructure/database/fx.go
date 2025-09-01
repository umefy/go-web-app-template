package database

import (
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm"
	"go.uber.org/fx"
)

var Module = fx.Module("database",
	gorm.Module,
)
