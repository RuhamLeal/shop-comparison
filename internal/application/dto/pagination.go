package dto

type PaginatorInput struct {
	Skip  int64 `mapstructure:"skip"`
	Limit int64 `mapstructure:"limit"`
}

type PaginatorOutput struct {
	Total int64 `json:"total"`
}
