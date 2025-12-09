package entity

type PaginatorInput struct {
	Skip  int64
	Limit int64
}

type PaginatorOutput struct {
	Total int64
}
