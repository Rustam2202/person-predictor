package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Conn *sql.DB
}

func MustConnectToPostgres(ctx context.Context, cfg *Config) *Postgres {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		panic("failed connection to database")
	}
	return &Postgres{Conn: db}
}
