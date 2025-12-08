package entity

type Specification struct {
	ID      int64
	Title   string
	Content string
}

type SpecificationProps struct {
	ID      int64
	Title   string
	Content string
}

func NewSpecification(props SpecificationProps) Specification {
	return Specification{
		ID:      props.ID,
		Title:   props.Title,
		Content: props.Content,
	}
}
