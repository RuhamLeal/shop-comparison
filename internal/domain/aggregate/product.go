package aggregate

import "project/internal/domain/entity"

type ProductWithSpecificationsGroups struct {
	Product              *entity.Product
	SpecificationsGroups []*entity.SpecificationGroup
}
