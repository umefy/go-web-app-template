package http

import (
	"github.com/umefy/go-web-app-template/internal/delivery/graphql"
	v1 "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1"
	"go.uber.org/fx"
)

var Module = fx.Module("httpServer",
	fx.Provide(NewServer),
	v1.Module,
	graphql.Module,
)
