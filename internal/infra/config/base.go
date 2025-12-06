package config

import (
	"log"
	"path/filepath"
	"project/internal/infra/config/environment"

	"github.com/joho/godotenv"
)

type BaseConfig struct {
	Fiber  *environment.Fiber
	Sqlite *environment.Sqlite
}

func NewBaseConfig(envFilePath string) *BaseConfig {
	absolutePath, err := filepath.Abs(envFilePath)
	if err != nil {
		log.Fatalf("can't resolve absolute path. error: %v", err)
	}

	err = godotenv.Load(absolutePath)
	if err != nil {
		log.Fatalf("can't load .env file. error: %v", err)
	}

	return &BaseConfig{
		Fiber:  environment.NewFiberConfig(),
		Sqlite: environment.NewSqliteConfig(),
	}
}
