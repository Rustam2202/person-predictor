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

type mockRepository struct {
	repo PersonRepository
	mock sqlmock.Sqlmock
	ctx  context.Context
}

func NewMockRepository() mockRepository {
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
	return mockRepository{repo: *repo, mock: mock, ctx: ctx}
}

func TestCreate(t *testing.T) {
	var err error
	mr := NewMockRepository()

	mr.mock.ExpectBegin()
	mr.mock.ExpectQuery(`INSERT INTO "persons" (.+) VALUES (.+) RETURNING "id"`).
		WithArgs("John", "Doe", "", 0, "", "").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mr.mock.ExpectCommit()
	err = mr.repo.Create(mr.ctx, &domain.Person{Name: "John", Surname: "Doe"})
	assert.Nil(t, err)

	mr.mock.ExpectBegin()
	mr.mock.ExpectQuery(`INSERT INTO "persons" (.+) VALUES (.+) RETURNING "id"`).
		WithArgs("", "", "", 0, "", "").
		WillReturnError(assert.AnError)
	mr.mock.ExpectRollback()
	err = mr.repo.Create(mr.ctx, &domain.Person{})
	assert.Error(t, err)
}

func TestGet(t *testing.T) {
	mr := NewMockRepository()
	{
		filters := make(map[string]interface{})
		filters["name"] = "John"
		filters["surname"] = "Doe"
		mr.mock.ExpectQuery(`SELECT (.+) FROM "persons" WHERE "name" = (.+) AND "surname" = (.+)`).
			WithArgs("John", "Doe").
			WillReturnRows(sqlmock.NewRows(
				[]string{"Id", "Name", "Surname", "Patronymic", "Age", "Gender", "Country"}).
				AddRow(int64(1), "John", "Doe", "", "42", "male", "US").
				AddRow(int64(2), "John", "Doe", "", "24", "male", "CA"),
			)
		persons, err := mr.repo.Get(mr.ctx, filters, 0)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(persons))
	}
	{
		filters := make(map[string]interface{})

		mr.mock.ExpectQuery(`SELECT (.+) FROM "persons"`).
			WithArgs().
			WillReturnError(assert.AnError)
		persons, err := mr.repo.Get(mr.ctx, filters, 0)
		assert.Error(t, err)
		assert.Nil(t, persons)
	}
}

func TestDelete(t *testing.T) {
	var err error
	mr := NewMockRepository()

	mr.mock.ExpectBegin()
	mr.mock.ExpectExec(`DELETE FROM "persons" WHERE "persons"."id" = (.+)`).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	mr.mock.ExpectCommit()
	err = mr.repo.Delete(mr.ctx, 1)
	assert.Nil(t, err)

	mr.mock.ExpectBegin()
	mr.mock.ExpectExec(`DELETE FROM "persons" WHERE "persons"."id" = (.+)`).
		WithArgs(99).WillReturnResult(driver.ResultNoRows)
	mr.mock.ExpectCommit()
	err = mr.repo.Delete(mr.ctx, 99)
	assert.Nil(t, err)
}
