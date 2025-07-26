package user

import (
	"net/http"

	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/handler/user/mapping"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
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

	resp := api.UserGetAllResponse{
		Data: sliceskit.Map(users, mapping.UserModelToApiUser),
	}

	return jsonkit.JSONResponse(w, http.StatusOK, &resp)
}
