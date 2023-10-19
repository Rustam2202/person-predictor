package database

import (
	"fmt"
	"person-predicator/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GORM struct {
	Gorm *gorm.DB
}

func MustConnectToGormPostgres(cfg *Config) *GORM {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed connection to database")
	}
	err = db.AutoMigrate(&domain.Person{})
	if err != nil {
		panic("Failed to migrate table")
	}
	return &GORM{Gorm: db}
}
