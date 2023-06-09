package httptools

import (
	"net/url"
	"strconv"
)

var (
	paginationParamLimit  = "limit"
	paginationParamOffset = "offset"

	defaultPaginationLimit  = 20
	defaultPaginationOffset = 0

	paginationLimitMax = 1000
)

type PaginationConfig struct {
	ParamLimit   string
	ParamOffset  string
	DefaultLimit int
	LimitMax     int
}

func ConfigPagination(cfg *PaginationConfig) {
	paginationParamLimit = cfg.ParamLimit
	paginationParamOffset = cfg.ParamOffset
	defaultPaginationLimit = cfg.DefaultLimit
	paginationLimitMax = cfg.LimitMax
}

type Pagination struct {
	Limit  int
	Offset int
}

func (p *Pagination) params() map[string]string {
	return map[string]string{
		paginationParamLimit:  strconv.Itoa(p.Limit),
		paginationParamOffset: strconv.Itoa(p.Offset),
	}
}

func ParsePaginationFromURL(u *url.URL) *Pagination {
	p := &Pagination{
		Limit:  defaultPaginationLimit,
		Offset: defaultPaginationOffset,
	}

	q := u.Query()

	val := q.Get(paginationParamLimit)
	if val != "" {
		limit, err := strconv.Atoi(val)
		if err != nil || (limit < 1 || limit > paginationLimitMax) {
			limit = defaultPaginationLimit
		}
		p.Limit = limit
	}

	val = q.Get(paginationParamOffset)
	if val != "" {
		offset, err := strconv.Atoi(val)
		if err != nil {
			p.Offset = offset
		}
	}

	return p
}

func ParsePaginationFromURLString(rawURL string) (*Pagination, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return ParsePaginationFromURL(u), nil
}

func GenerateNextPagignatedURL(u *url.URL) *url.URL {
	p := ParsePaginationFromURL(u)
	p.Offset += p.Limit
	return AddParamsToURL(u, p.params())
}
