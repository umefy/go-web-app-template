package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/server/httpserver/user/mapping"
	api "github.com/umefy/go-web-app-template/openapi/protogen/v1/models"
	"github.com/umefy/go-web-app-template/pkg/validation"
	"github.com/umefy/godash/jsonkit"
)

type UpdateUserInput struct {
	api.UpdateUserInput
}

var _ validation.Validate = (*UpdateUserInput)(nil)

func (u *UpdateUserInput) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Age,
			validation.MinWrapperspb(12),
			validation.MaxWrapperspb(20),
		),
	)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	h.loggerService.DebugContext(ctx, "UpdateUser")

	var input UpdateUserInput
	if err := jsonkit.BindProtoRequestBody(r, &input); err != nil {
		return err
	}

	if err := input.Validate(); err != nil {
		return err
	}

	user := mapping.ApiUpdateUserInputToUserModel(&input.UpdateUserInput)

	userID := r.PathValue("id")
	user, err := h.userService.UpdateUser(ctx, userID, user)
	if err != nil {
		return err
	}

	userResponse := mapping.UserModelToApiUser(user)

	return jsonkit.ProtoJSONResponse(w, http.StatusOK, userResponse)
}
