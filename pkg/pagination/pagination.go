package pagination

import "strconv"

const (
	DefaultOffset       = 0
	DefaultPageSize     = 25
	DefaultIncludeTotal = false
)

type Option = func(p *Pagination)

type Pagination struct {
	Offset       int
	PageSize     int
	IncludeTotal bool
}

func NewPagination(offset, pagesize, includeTotal string, options ...Option) Pagination {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = DefaultOffset
	}
	pagesizeInt, err := strconv.Atoi(pagesize)
	if err != nil {
		pagesizeInt = DefaultPageSize
	}

	var p = Pagination{
		Offset:       offsetInt,
		PageSize:     pagesizeInt,
		IncludeTotal: includeTotal == "true",
	}

	for _, option := range options {
		option(&p)
	}

	if p.Offset < 0 {
		p.Offset = DefaultOffset
	}

	if p.PageSize <= 0 {
		p.PageSize = DefaultPageSize
	}

	return p
}

func WithDefaultOffset(offset int) Option {
	return func(p *Pagination) {
		if p.Offset < 0 {
			p.Offset = offset
		}
	}
}

func WithDefaultPageSize(pageSize int) Option {
	return func(p *Pagination) {
		if p.PageSize <= 0 {
			p.PageSize = pageSize
		}
	}
}

func WithDefaultIncludeTotal(includeTotal bool) Option {
	return func(p *Pagination) {
		if !p.IncludeTotal {
			p.IncludeTotal = includeTotal
		}
	}
}
