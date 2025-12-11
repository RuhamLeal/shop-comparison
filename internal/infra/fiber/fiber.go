package fiber

import (
	"fmt"
	"project/internal/infra/config/environment"
	"project/internal/infra/fiber/route"
	"project/internal/infra/fiber/utils/response"
	"project/internal/infra/sqlite"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type Fiber struct {
	prefork bool
	address string
	App     *fiber.App
}

func NewFiberInstance(
	config *environment.Fiber,
	sqlite *sqlite.Sqlite,
) *Fiber {
	app := fiber.New(
		fiber.Config{
			AppName:     "Edge DNS API",
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
			ErrorHandler: func(c fiber.Ctx, err error) error {
				log.Error(err)
				if c.Response().StatusCode() == fiber.StatusNotFound {
					return response.SendNotFound(c, "sorry, endpoint not found", err)
				}

				return response.SendInternalServerError(c, "sorry, unexpected error occurred", err)
			},
		},
	)

	router := route.NewRouter(app, sqlite)

	router.Load()

	app.Use(
		func(c fiber.Ctx) error {
			return response.SendNotFound(c, "Sorry, endpoint not found - unknown route")
		},
	)

	return &Fiber{App: app, prefork: config.Prefork, address: fmt.Sprintf("%s:%d", config.Host, config.Port)}
}

func (f *Fiber) Start() {
	listenErr := f.App.Listen(
		f.address,
		fiber.ListenConfig{
			EnablePrefork: f.prefork,
		})

	if listenErr != nil {
		panic(fmt.Errorf("oops... server is not running! error: %v", listenErr))
	}
}
