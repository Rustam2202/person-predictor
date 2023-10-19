package main

import (
	"context"
	"person-predicator/internal/config"
	"person-predicator/internal/database"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoadConfig()
	database.MustConnectToPostgres(ctx, cfg.Database)
}
