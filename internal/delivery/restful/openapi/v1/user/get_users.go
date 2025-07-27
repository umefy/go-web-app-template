package user

import (
	"net/http"

	api "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/generated"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/user/mapping"
	"github.com/umefy/godash/jsonkit"
	"github.com/umefy/godash/sliceskit"
)

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	h.logger.DebugContext(ctx, "GetUsers")

	users, err := h.userService.GetUsers(ctx)
	if err != nil {
		return err
	}

	resp := api.UserGetAllResponse{
		Data: sliceskit.Map(users, mapping.UserModelToApiUser),
	}

	return jsonkit.JSONResponse(w, http.StatusOK, &resp)
}
