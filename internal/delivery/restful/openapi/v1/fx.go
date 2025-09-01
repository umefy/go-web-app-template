package v1

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/user"
	"go.uber.org/fx"
)

const (
	FX_TAG_NAME_API_V1_ROUTER   = `name:"apiV1Router"`
	FX_TAG_GROUP_API_V1_ROUTERS = `group:"apiV1Routers"`
)

var Module = fx.Module("apiV1Router",
	fx.Provide(
		fx.Annotate(
			NewApiV1Router,
			fx.As(new(http.Handler)),
			fx.ResultTags(FX_TAG_NAME_API_V1_ROUTER),
		),
		fx.Annotate(
			user.NewHandler,
			fx.As(new(handler.Router)),
			fx.ResultTags(FX_TAG_GROUP_API_V1_ROUTERS),
		),
	),
)
