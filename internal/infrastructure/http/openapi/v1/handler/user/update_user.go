package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/infrastructure/http/handler/middleware"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/openapi/v1/handler/user/mapping"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
	"github.com/umefy/godash/jsonkit"
)

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var input api.UserUpdate
	if err := jsonkit.BindRequestBody(r, &input); err != nil {
		return err
	}

	userUpdateInput := mapping.ApiUserUpdateToUserModelUpdate(&input)

	userID := r.PathValue("id")

	tx, err := middleware.GetTransaction(ctx)
	if err != nil {
		return err
	}
	user, err := h.userService.UpdateUser(ctx, userID, userUpdateInput, tx)
	if err != nil {
		return err
	}

	userResp := mapping.UserModelToApiUser(user)
	resp := api.UserUpdateResponse{
		Data: &userResp,
	}

	return jsonkit.JSONResponse(w, http.StatusOK, &resp)
}
