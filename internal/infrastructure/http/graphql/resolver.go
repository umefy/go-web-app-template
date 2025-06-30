package graphql

import (
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	loggerSvc "github.com/umefy/go-web-app-template/internal/domain/logger/service"
	userSvc "github.com/umefy/go-web-app-template/internal/domain/user/service"
)

type Resolver struct {
	UserService   userSvc.Service
	LoggerService loggerSvc.Service
	DbQuery       *query.Query
}
