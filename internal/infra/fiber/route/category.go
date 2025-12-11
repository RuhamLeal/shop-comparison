package route

import (
	"project/internal/infra/fiber/handler"

	"github.com/gofiber/fiber/v3"
)

func (r *Router) loadCategoryRoutes(router fiber.Router) {
	handler := handler.NewCategory(r.Sqlite)

	router.Get("/categories",
		handler.GetAllCategoriesHandler,
	)
}
