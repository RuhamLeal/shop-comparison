package entity

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       int64 // in cents R$ 5.012,00 -> 5012
	Rating      int8  // 0-50 (10 = 1 star, 25 = 2.5 stars, 50 = 5 stars)
}

type ProductProps struct {
	ID          int64
	Name        string
	Description string
	Price       int64
	Rating      int8
}

func NewProduct(props ProductProps) Product {
	return Product{
		ID:          props.ID,
		Name:        props.Name,
		Description: props.Description,
		Price:       props.Price,
		Rating:      props.Rating,
	}
}
