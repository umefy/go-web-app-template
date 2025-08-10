package v1

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/user"
	"go.uber.org/fx"
)

var Module = fx.Module("apiV1Router",
	fx.Provide(
		fx.Annotate(
			NewApiV1Router,
			fx.As(new(http.Handler)),
			fx.ResultTags(`name:"apiV1Router"`),
		),
		fx.Annotate(
			user.NewHandler,
			fx.As(new(handler.Router)),
			fx.ResultTags(`group:"apiV1Routers"`),
		),
	),
)
