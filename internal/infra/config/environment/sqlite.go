package environment

import (
	"fmt"
	"project/internal/infra/config/services"
	"strconv"
)

type Sqlite struct {
	Path        string
	BusyTimeout int64
	Dsn         string
}

func NewSqliteConfig() *Sqlite {
	timeout := services.GetEnvironmentVariable("SQLITE_BUSY_TIMEOUT", false)

	timeoutInt, err := strconv.ParseInt(timeout, 10, 64)

	if err != nil {
		panic(fmt.Sprintf("Invalid value for 'SQLITE_BUSY_TIMEOUT' env, value: %s", timeout))
	}

	path := services.GetEnvironmentVariable("SQLITE_PATH", false)

	return &Sqlite{
		BusyTimeout: timeoutInt,
		Path:        path,
		Dsn: fmt.Sprintf(
			"file:%s?_busy_timeout=%d&_fk=1", path, timeoutInt,
		),
	}
}
