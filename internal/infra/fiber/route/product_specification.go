package route

import (
	"project/internal/application/dto"
	"project/internal/infra/fiber/handler"
	"project/internal/infra/fiber/middleware"
	"project/internal/infra/fiber/schemas"

	"github.com/gofiber/fiber/v3"
)

func (r *Router) loadProductSpecificationRoutes(router fiber.Router) {
	handler := handler.NewProductSpecification(r.Sqlite)

	router.Get("/products/specifications",
		handler.CreateOneProductSpecificationValueHandler,
		middleware.Validate[dto.CreateOneProductSpecificationValueInput](schemas.CreateOneProductSpecificationValueSchema),
	)

}
