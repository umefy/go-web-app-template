package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/server/httpserver/user/mapping"
	api "github.com/umefy/go-web-app-template/openapi/protogen/v1/models"
	"github.com/umefy/godash/jsonkit"
	"github.com/umefy/godash/sliceskit"
)

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	h.loggerService.DebugContext(ctx, "GetUsers")

	users, err := h.userService.GetUsers(ctx)
	if err != nil {
		return err
	}

	usersResponse := api.GetUsersResponse{
		Data: sliceskit.Map(users, mapping.UserModelToApiUser),
	}

	return jsonkit.ProtoJSONResponse(w, http.StatusOK, &usersResponse)
}
