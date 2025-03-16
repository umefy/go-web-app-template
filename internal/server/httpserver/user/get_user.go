package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/server/httpserver/user/mapping"
	"github.com/umefy/godash/jsonkit"
)

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	h.loggerService.DebugContext(ctx, "GetUser")

	userID := r.PathValue("id")

	user, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	userResponse := mapping.UserModelToApiUser(user)

	return jsonkit.JSONResponse(w, http.StatusOK, userResponse)
}
