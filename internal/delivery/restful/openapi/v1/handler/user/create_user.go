package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/restful/handler/middleware"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/handler/user/mapping"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
	"github.com/umefy/godash/jsonkit"
)

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	h.loggerService.DebugContext(ctx, "CreateUser")

	var userInput api.UserCreate
	if err := jsonkit.BindRequestBody(r, &userInput); err != nil {
		return err
	}

	userCreateInput := mapping.ApiUserCreateToUserModelCreate(&userInput)

	tx, err := middleware.GetTransaction(ctx)
	if err != nil {
		return err
	}

	user, err := h.userService.CreateUser(ctx, userCreateInput, tx)
	if err != nil {
		return err
	}

	userResp := mapping.UserModelToApiUser(user)
	resp := api.UserCreateResponse{
		Data: &userResp,
	}

	return jsonkit.JSONResponse(w, http.StatusOK, &resp)
}
