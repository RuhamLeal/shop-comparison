package middleware

import (
	"fmt"
	"log"
	exceptions "project/internal/domain/exception"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Default struct{}

func NewDefault() *Default {
	return &Default{}
}

func (m *Default) Recoverer() fiber.Handler {
	return func(c fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error)

				if !ok {
					err = fmt.Errorf("%v", r)
				}

				log.Print(err)

				err = exceptions.Usecase(err, exceptions.UsecaseOpts{
					Code:        "Recoverer",
					StatusCode:  fiber.StatusInternalServerError,
					Message:     "Unexpected Error",
					StackLength: 10,
				})
			}
		}()

		return c.Next()
	}
}

func (m *Default) Cors() fiber.Handler {
	return cors.New()
}

func (m *Default) Logger() fiber.Handler {
	return logger.New()
}
