package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/infrastructure/http/handler/user/mapping"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/middleware"
	api "github.com/umefy/go-web-app-template/openapi/protogen/v1/models"
	"github.com/umefy/go-web-app-template/pkg/validation"
	"github.com/umefy/godash/jsonkit"
)

type CreateUserInput struct {
	api.UserCreate
}

var _ validation.Validate = (*CreateUserInput)(nil)

func (u *CreateUserInput) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Age,
			validation.Required.Error("must be provided"),
			validation.Min(12),
			validation.Max(20),
		),
	)
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	h.loggerService.DebugContext(ctx, "CreateUser")

	var userInput CreateUserInput
	if err := jsonkit.BindProtoRequestBody(r, &userInput); err != nil {
		return err
	}

	if err := (&userInput).Validate(); err != nil {
		return err
	}

	user := mapping.ApiUserCreateToUserModel(&userInput.UserCreate)

	tx, err := middleware.GetTransaction(ctx)
	if err != nil {
		return err
	}

	user, err = h.userService.CreateUser(ctx, user, tx)
	if err != nil {
		return err
	}

	resp := api.UserCreateResponse{
		Data: mapping.UserModelToApiUser(user),
	}

	return jsonkit.ProtoJSONResponse(w, http.StatusOK, &resp)
}
