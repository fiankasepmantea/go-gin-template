package base

type PaginationResponse[T any] struct {
	TotalData   int64 `json:"total_data"`
	TotalPages  int   `json:"total_pages"`
	CurrPage    int   `json:"curr_page"`
	NextPage    int   `json:"next_page"`
	PrevPage    int   `json:"prev_page,omitempty"`
	DataPerPage int   `json:"data_per_page"`
	Data        []T   `json:"data"`
}

func BuildPagination[T any](page, limit int, total int64, data []T) PaginationResponse[T] {
	totalPages := 0
	if limit > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	nextPage := page
	if page < totalPages {
		nextPage = page + 1
	}

	resp := PaginationResponse[T]{
		TotalData:   total,
		TotalPages:  totalPages,
		CurrPage:    page,
		NextPage:    nextPage,
		DataPerPage: limit,
		Data:        data,
	}

	if page > 1 {
		resp.PrevPage = page - 1
	}

	return resp
}