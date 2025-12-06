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
	host := services.GetEnvironmentVariable("FIBER_HOST", false)
	portEnv := services.GetEnvironmentVariable("FIBER_PORT", false)

	port, err := strconv.Atoi(portEnv)

	if err != nil {
		panic(fmt.Sprintf("Invalid value for 'FIBER_PORT' env, value: %s", portEnv))
	}

	debug := services.GetEnvironmentVariableAsBool("FIBER_DEBUG", false)

	prefork := services.GetEnvironmentVariableAsBool("FIBER_PREFORK", false)

	return &Fiber{
		Host:    host,
		Port:    port,
		Debug:   debug,
		Prefork: prefork,
	}
}
