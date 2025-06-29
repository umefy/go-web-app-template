package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/infrastructure/http/openapi/v1/handler/user/mapping"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
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
