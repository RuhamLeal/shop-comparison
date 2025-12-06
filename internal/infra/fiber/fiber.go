package fiber

import (
	"fmt"
	"project/internal/infra/config/environment"
	"project/internal/infra/sqlite"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
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
