package route

import (
	"log"
	"project/internal/infra/config/services"
	"project/internal/infra/fiber/middleware"
	"project/internal/infra/sqlite"

	"github.com/Flussen/swagger-fiber-v3"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/basicauth"
)

type Router struct {
	App    *fiber.App
	Sqlite *sqlite.Sqlite
}

func NewRouter(
	app *fiber.App,
	sqlite *sqlite.Sqlite,
) *Router {
	return &Router{
		App:    app,
		Sqlite: sqlite,
	}
}

func (r *Router) Load() {
	r.registerMiddlewares()
	r.registerRoutes()
}

func (r *Router) registerMiddlewares() {
	r.loadDefaultMiddlewares()
}

func (r *Router) registerRoutes() {
	r.loadSwaggerRoutes()
	r.loadMainRoutes()
}

func (r *Router) loadDefaultMiddlewares() {
	defaultMiddleware := middleware.NewDefault()

	r.App.Use(
		defaultMiddleware.Recoverer(),
		defaultMiddleware.Cors(),
		defaultMiddleware.Logger(),
	)
}

func (r *Router) loadSwaggerRoutes() {
	log.Print(services.GetEnvironmentVariable("SWAGGER_ROUTE_ACCESS_USER", false), services.GetEnvironmentVariable("SWAGGER_ROUTE_ACCESS_PASSWORD", false))
	swaggerGroup := r.App.Group("/swagger")
	swaggerGroup.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			services.GetEnvironmentVariable("SWAGGER_ROUTE_ACCESS_USER", false): services.GetEnvironmentVariable("SWAGGER_ROUTE_ACCESS_PASSWORD", false),
		},
	}))
	swaggerGroup.Get("*", swagger.New())
}

func (r *Router) loadMainRoutes() {
	privateGroup := r.App.Group("/")

	r.loadCategoryRoutes(privateGroup)
	r.loadProductRoutes(privateGroup)
	r.loadProductSpecificationRoutes(privateGroup)
	r.loadSpecificationRoutes(privateGroup)
	r.loadSpecificationGroupRoutes(privateGroup)
}
