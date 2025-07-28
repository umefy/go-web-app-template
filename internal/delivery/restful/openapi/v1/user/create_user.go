package user

import (
	"net/http"

	api "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/generated"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/user/mapping"
	"github.com/umefy/godash/jsonkit"
)

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	h.logger.DebugContext(ctx, "CreateUser")

	var userInput api.UserCreate
	if err := jsonkit.BindRequestBody(r, &userInput); err != nil {
		return err
	}

	userCreateInput := mapping.ApiUserCreateToUserModelCreate(&userInput)

	user, err := h.userService.CreateUser(ctx, userCreateInput)
	if err != nil {
		return err
	}

	userResp := mapping.UserModelToApiUser(user)
	resp := api.UserCreateResponse{
		Data: &userResp,
	}

	return jsonkit.JSONResponse(w, http.StatusOK, &resp)
}
