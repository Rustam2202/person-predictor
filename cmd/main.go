package main

import (
	"context"
	"os"
	"os/signal"
	"person-predicator/internal/config"
	"person-predicator/internal/database"
	"person-predicator/internal/logger"
	"person-predicator/internal/repository"
	"person-predicator/internal/server"
	"person-predicator/internal/server/handlers/persons"
	"person-predicator/internal/service"
	"sync"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	wg := &sync.WaitGroup{}

	cfg := config.MustLoadConfig()
	logger.MustConfigLogger(cfg.Logger)
	db := database.MustConnectToGormPostgres(cfg.Database)
	personRepository := repository.NewPersonRepository(db)
	personService := service.NewPersonService(personRepository)
	personHandler := persons.NewPersonHandler(personService)
	s := server.NewHTTP(cfg.Server, personHandler)
	wg.Add(1)
	go s.StartHTTP(ctx, wg)
	wg.Wait()
}
