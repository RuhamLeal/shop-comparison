package route

import (
	"project/internal/infra/fiber/handler"

	"github.com/gofiber/fiber/v3"
)

func (r *Router) loadSpecificationGroupRoutes(router fiber.Router) {
	handler := handler.NewSpecificationGroup(r.Sqlite)

	router.Get("/specifications/groups",
		handler.GetAllSpecificationGroupsHandler,
	)
}
