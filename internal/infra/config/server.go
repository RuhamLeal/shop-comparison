package config

import (
	"project/internal/infra/fiber"
	"project/internal/infra/sqlite"
)

type Server struct {
	Fiber  *fiber.Fiber
	Sqlite *sqlite.Sqlite
}

func NewServerInstances(config *BaseConfig) *Server {
	// order is relevant

	sqlite := sqlite.NewSqliteInstance(config.Sqlite)

	fiber := fiber.NewFiberInstance(config.Fiber, sqlite)

	return &Server{
		Fiber: fiber,
	}
}

func (s *Server) Start() {
	s.Fiber.Start()
}

func (s *Server) Stop() {
	s.Fiber.App.Shutdown()
}
