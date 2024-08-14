package paginate

type PaginateRes[T any] struct {
	Data     T            `json:"data"`
	Paginate PaginateMeta `json:"paginate"`
}
