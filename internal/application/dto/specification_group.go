package dto

type GetAllSpecificationGroupsOutput struct {
	Groups []*SpecificationGroupOutput `json:"groups"`
}

type SpecificationGroupOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
