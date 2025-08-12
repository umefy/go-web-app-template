package mapping

import (
	api "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/generated"
	"github.com/umefy/go-web-app-template/pkg/cast"
	"github.com/umefy/go-web-app-template/pkg/pagination"
)

func PaginationMetadataToApiPaginationMetadata(metadata *pagination.PaginationMetadata) *api.PaginationMetadata {

	return &api.PaginationMetadata{
		Offset:   metadata.Offset,
		PageSize: metadata.PageSize,
		Count:    metadata.Count,
		HasMore:  metadata.HasMore,
		Total:    cast.Int64PtrToIntPtr(metadata.Total),
	}
}
