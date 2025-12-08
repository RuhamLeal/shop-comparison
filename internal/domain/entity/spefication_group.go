package entity

type SpecificationGroup struct {
	ID                int64
	Name              string
	Description       string
	TotalSpefications int64
	Specifications    []*Specification
}

type SpecificationGroupProps struct {
	ID                int64
	Name              string
	Description       string
	TotalSpefications int64
	Specifications    []*Specification
}

func NewSpecificationGroup(props SpecificationGroupProps) SpecificationGroup {
	return SpecificationGroup{
		ID:                props.ID,
		Name:              props.Name,
		Description:       props.Description,
		TotalSpefications: props.TotalSpefications,
		Specifications:    props.Specifications,
	}
}

func (sg *SpecificationGroup) HasSpecifications() bool {
	return len(sg.Specifications) > 0
}
