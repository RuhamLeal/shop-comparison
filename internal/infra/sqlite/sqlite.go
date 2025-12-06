package sqlite

import (
	"database/sql"
	"fmt"
	"project/internal/infra/config/environment"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	DB *sql.DB
}

func NewSqliteInstance(config *environment.Sqlite) *Sqlite {
	dbConn, err := sql.Open("sqlite3", config.Dsn)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to sqlite database, err: %v", err))
	}

	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		panic(fmt.Sprintf("Error pinging to sqlite database, err: %v", err))
	}

	return &Sqlite{DB: dbConn}
}
