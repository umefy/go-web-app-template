package graphql

import (
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	userSvc "github.com/umefy/go-web-app-template/internal/domain/user/service"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
)

type Resolver struct {
	UserService userSvc.Service
	Logger      logger.Logger
	DbQuery     *query.Query
}
