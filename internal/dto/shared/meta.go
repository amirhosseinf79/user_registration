package shared

type FieldPagination struct {
	Page     int `query:"page"`
	PageSize int `query:"pageSize"`
}

type MetaPagination struct {
	TotalCount  int `json:"totalCount"`
	PageSize    int `json:"pageSize"`
	CurrentPage int `json:"currentPage"`
	NextPage    int `json:"nextPage"`
	EndPage     int `json:"endPage"`
}
