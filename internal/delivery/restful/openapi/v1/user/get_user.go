package user

import (
	"net/http"

	api "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/generated"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/mapping"
	"github.com/umefy/godash/jsonkit"
)

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	userID := r.PathValue("id")

	user, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	userResp := mapping.UserModelToApiUser(user)
	resp := api.UserGetResponse{
		Data: &userResp,
	}

	return jsonkit.JSONResponse(w, http.StatusOK, &resp)
}
