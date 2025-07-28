package user

import (
	"net/http"

	api "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/generated"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/user/mapping"
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

	user, err := h.userService.UpdateUser(ctx, userID, userUpdateInput)
	if err != nil {
		return err
	}

	userResp := mapping.UserModelToApiUser(user)
	resp := api.UserUpdateResponse{
		Data: &userResp,
	}

	return jsonkit.JSONResponse(w, http.StatusOK, &resp)
}
