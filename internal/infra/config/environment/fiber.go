package environment

import (
	"fmt"
	"project/internal/infra/config/services"
	"strconv"
)

type Fiber struct {
	Host    string
	Port    int
	Debug   bool
	Prefork bool
}

func NewFiberConfig() *Fiber {
	host := services.GetEnvironmentVariable("HOST", false)
	portEnv := services.GetEnvironmentVariable("APP_PORT", false)

	port, err := strconv.Atoi(portEnv)

	if err != nil {
		panic(fmt.Sprintf("Invalid value for 'APP_PORT' env, value: %s", portEnv))
	}

	appDebugEnv := services.GetEnvironmentVariable("APP_DEBUG", false)

	debug, err := strconv.ParseBool(appDebugEnv)

	if err != nil {
		panic(fmt.Sprintf("Invalid value for 'APP_DEBUG' env, value: %s", appDebugEnv))
	}

	preforkEnv := services.GetEnvironmentVariable("FIBER_PREFORK", false)

	prefork, err := strconv.ParseBool(preforkEnv)

	if err != nil {
		panic(fmt.Sprintf("invalid 'FIBER_PREFORK' env value: %s", preforkEnv))
	}

	return &Fiber{
		Host:    host,
		Port:    port,
		Debug:   debug,
		Prefork: prefork,
	}
}
