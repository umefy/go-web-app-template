package graphql

import (
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	userSvc "github.com/umefy/go-web-app-template/internal/service/user"
)

type Resolver struct {
	UserService userSvc.Service
	Logger      logger.Logger
	DbQuery     *query.Query
}
