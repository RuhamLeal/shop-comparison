package dto

type PaginatorInput struct {
	Skip  int64 `json:"skip" mapstructure:"skip"`
	Limit int64 `json:"limit" mapstructure:"limit"`
}

type PaginatorOutput struct {
	Total int64 `json:"total" mapstructure:"total"`
}
