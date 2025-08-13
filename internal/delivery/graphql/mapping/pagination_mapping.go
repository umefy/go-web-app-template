package mapping

import (
	"github.com/umefy/go-web-app-template/internal/delivery/graphql/model"
	"github.com/umefy/go-web-app-template/pkg/cast"
	"github.com/umefy/go-web-app-template/pkg/pagination"
)

func PaginationMetadataToGraphqlPaginationMetadata(metadata *pagination.PaginationMetadata) *model.PaginationMetadata {

	if metadata == nil {
		return nil
	}

	return &model.PaginationMetadata{
		Offset:   int32(metadata.Offset),
		PageSize: int32(metadata.PageSize),
		Count:    int32(metadata.Count),
		HasMore:  metadata.HasMore,
		Total:    cast.Int64PtrToIntPtr(metadata.Total),
	}
}
