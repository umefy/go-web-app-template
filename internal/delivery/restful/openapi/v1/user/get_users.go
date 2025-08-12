package user

import (
	"net/http"

	api "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/generated"
	apiMapping "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/mapping"
	userMapping "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/user/mapping"
	"github.com/umefy/go-web-app-template/pkg/pagination"
	"github.com/umefy/godash/jsonkit"
	"github.com/umefy/godash/sliceskit"
)

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	h.logger.DebugContext(ctx, "GetUsers")

	query := r.URL.Query()

	users, paginationMetadata, err := h.userService.GetUsers(ctx, pagination.NewPagination(query.Get("offset"), query.Get("pageSize"), query.Get("includeTotal")))
	if err != nil {
		return err
	}

	resp := api.UserGetAllResponse{
		Data:     sliceskit.Map(users, userMapping.UserModelToApiUser),
		PageInfo: apiMapping.PaginationMetadataToApiPaginationMetadata(paginationMetadata),
	}

	return jsonkit.JSONResponse(w, http.StatusOK, &resp)
}
