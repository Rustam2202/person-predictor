package repository

import (
	"context"
	"database/sql/driver"
	"person-predicator/internal/database"
	"person-predicator/internal/domain"
	"person-predicator/internal/logger"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	db, mock, _ := sqlmock.New()

	gormDB, _ := gorm.Open(
		postgres.New(postgres.Config{DriverName: "postgres", Conn: db}))
	repo := NewPersonRepository(&database.GORM{Gorm: gormDB})

	// mock.ExpectQuery("DELETE FROM persons").WithArgs(1).WillReturnError(nil)
	mock.ExpectBegin()

	// mock.ExpectQuery("DELETE FROM \"person\" WHERE \"person\".\"id\" = $1").WithArgs(1,1).WillReturnError(nil)

	mock.ExpectExec("INSERT INTO persons (name, surname) VALUES ($1,$2) RETURNING id").
		WithArgs("John", "Doe").
		WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectRollback()
	// 		mock.ExpectCommit()

	repo.Create(ctx, &domain.Person{Name: "John", Surname: "Doe"})
}

func TestDelete(t *testing.T) {
	var err error
	ctx := context.Background()
	logger.MustConfigLogger(&logger.Config{
		Encoding:         "json",
		Level:            "debug",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	})
	db, mock, _ := sqlmock.New()

	gormDB, _ := gorm.Open(
		postgres.New(
			postgres.Config{DriverName: "postgres", Conn: db}),
		&gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	repo := NewPersonRepository(&database.GORM{Gorm: gormDB})

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"persons\" WHERE \"persons\".\"id\" = (.+)").
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	mock.ExpectCommit()
	err = repo.Delete(ctx, 1)
	assert.Error(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"persons\" WHERE \"persons\".\"id\" = (.+)").
		WithArgs(99).WillReturnResult(driver.ResultNoRows)
	mock.ExpectCommit()
	err = repo.Delete(ctx, 99)
	assert.Error(t, err)

}
