package route

import (
	"project/internal/application/dto"
	"project/internal/infra/fiber/handler"
	"project/internal/infra/fiber/middleware"
	"project/internal/infra/fiber/schemas"

	"github.com/gofiber/fiber/v3"
)

func (r *Router) loadProductRoutes(router fiber.Router) {
	handler := handler.NewProduct(r.Sqlite)

	router.Post("/products",
		handler.CreateOneProductHandler,
		middleware.Validate[dto.CreateOneProductInput](schemas.CreateOneProductSchema),
	)

	router.Post("/products/compare",
		handler.CompareProductsHandler,
		middleware.Validate[dto.CompareProductsInput](schemas.CompareProductsSchema),
	)

	router.Delete("/products/:public_id",
		handler.DeleteOneProductHandler,
		middleware.Validate[dto.DeleteOneProductInput](schemas.DeleteOneProductSchema),
	)

	router.Get("/categories/:category_public_id/products",
		handler.GetAllProductsByCategoryIdHandler,
		middleware.Validate[dto.GetAllProductsByCategoryIdInput](schemas.GetAllProductsByCategoryIdSchema),
	)

	router.Get("/products",
		handler.GetAllProductsHandler,
		middleware.Validate[dto.GetAllProductsInput](schemas.GetAllProductsSchema),
	)

	router.Get("/products/:public_id",
		handler.GetOneProductByPublicIdHandler,
		middleware.Validate[dto.GetOneProductByPublicIdInput](schemas.GetOneProductByPublicIdSchema),
	)

	router.Get("/products/:public_id/specifications",
		handler.GetOneProductWithSpecificationsByPublicIdHandler,
		middleware.Validate[dto.GetOneProductWithSpecificationsByPublicIdInput](schemas.GetOneProductWithSpecificationsByPublicIdSchema),
	)

	router.Put("/products/:public_id",
		handler.UpdateOneProductHandler,
		middleware.Validate[dto.UpdateOneProductInput](schemas.UpdateOneProductSchema),
	)
}
