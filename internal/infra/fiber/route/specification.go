package route

import (
	"project/internal/application/dto"
	"project/internal/infra/fiber/handler"
	"project/internal/infra/fiber/middleware"
	"project/internal/infra/fiber/schemas"

	"github.com/gofiber/fiber/v3"
)

func (r *Router) loadSpecificationRoutes(router fiber.Router) {
	handler := handler.NewSpecification(r.Sqlite)

	router.Get("/specifications",
		handler.GetAllSpecificationsHandler,
		middleware.Validate[dto.GetAllSpecificationsInput](schemas.GetAllSpecificationsSchema),
	)
}
