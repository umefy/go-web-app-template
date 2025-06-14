package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/server/httpserver/user/mapping"
	api "github.com/umefy/go-web-app-template/openapi/protogen/v1/models"
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

	resp := api.UserGetResponse{
		Data: mapping.UserModelToApiUser(user),
	}

	return jsonkit.ProtoJSONResponse(w, http.StatusOK, &resp)
}
