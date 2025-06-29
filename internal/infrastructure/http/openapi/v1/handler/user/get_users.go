package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/umefy/go-web-app-template/internal/infrastructure/http/openapi/v1/handler/user/mapping"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
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

	h.loggerService.DebugContext(ctx, "resp", slog.Any("resp", resp))
	// return jsonkit.JSONResponse(w, http.StatusOK, &resp)
	v, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.Write(v)
	return nil
}
