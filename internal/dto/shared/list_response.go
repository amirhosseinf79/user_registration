package shared

import "math"

type ResponseList[T any] struct {
	Items []T            `json:"items"`
	Meta  MetaPagination `json:"meta"`
}

func NewResponseList[T any](list []T, total, currentPage, pageSize int) ResponseList[T] {
	realPage := 1
	realPageSize := 10

	if currentPage != 0 {
		realPage = currentPage
	}
	if pageSize != 0 {
		realPageSize = pageSize
	}

	endPage := int(math.Ceil(float64(total) / float64(realPageSize)))
	nextPage := 0
	if realPage*realPageSize < total {
		nextPage = realPage + 1
	}

	meta := MetaPagination{
		TotalCount:  total,
		PageSize:    realPageSize,
		CurrentPage: realPage,
		NextPage:    nextPage,
		EndPage:     endPage,
	}

	return ResponseList[T]{
		Items: list,
		Meta:  meta,
	}
}
