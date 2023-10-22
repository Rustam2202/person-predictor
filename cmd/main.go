package main

import (
	"context"
	"person-predicator/internal/config"
	"person-predicator/internal/database"
	"person-predicator/internal/logger"
	"person-predicator/internal/repository"
	"person-predicator/internal/server"
	"person-predicator/internal/server/handlers/persons"
	"person-predicator/internal/service"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoadConfig()
	logger.MustConfigLogger(cfg.Logger)
	db := database.MustConnectToGormPostgres(cfg.Database)
	personRepository := repository.NewPersonRepository(db)
	personService := service.NewPersonService(personRepository)
	personHandler := persons.NewPersonHandler(personService)
	s := server.NewHTTP(cfg.Server, personHandler)
	s.StartHTTP(ctx)
}
