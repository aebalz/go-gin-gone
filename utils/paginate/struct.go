package paginate

type PaginateMeta struct {
	LastPage    int   `json:"last_page"`
	CurrentPage int   `json:"current_page"`
	Limit       int   `json:"limit"`
	Total       int64 `json:"total"`
}

type Param struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PaginateRes[T any] struct {
	Data     T            `json:"data"`
	Paginate PaginateMeta `json:"paginate"`
}
