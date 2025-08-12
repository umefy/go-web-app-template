package pagination

// PaginationMetadata is the metadata for the pagination returned for API to get more information about the pagination
type PaginationMetadata struct {
	Offset   int `json:"offset"`
	PageSize int `json:"pageSize"`
	// The number of items in the current page
	Count int `json:"count"`
	// Whether there are more items to fetch
	HasMore bool `json:"hasMore"`
	// The total number of items
	Total *int64 `json:"total,omitempty"`
}

func NewPaginationMetadata(offset, pageSize, count int, hasMore bool, total *int64) PaginationMetadata {
	return PaginationMetadata{
		Offset:   offset,
		PageSize: pageSize,
		Count:    count,
		HasMore:  hasMore,
		Total:    total,
	}
}
