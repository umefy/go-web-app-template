package graphql

import (
	loggerSvc "github.com/umefy/go-web-app-template/internal/domain/logger/service"
	userSvc "github.com/umefy/go-web-app-template/internal/domain/user/service"
)

type Resolver struct {
	UserService   userSvc.Service
	LoggerService loggerSvc.Service
}
