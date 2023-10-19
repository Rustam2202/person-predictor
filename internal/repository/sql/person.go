package repository

import (
	"person-predicator/internal/database"
	"person-predicator/internal/domain"
	"person-predicator/internal/logger"

	"go.uber.org/zap"
)

type PersonRepository struct {
	Db *database.Postgres
}

func NewPersonRepository(db *database.Postgres) *PersonRepository {
	return &PersonRepository{Db: db}
}

func (r *PersonRepository) Create(per *domain.Person) error {
	var lastInsertedId int64
	err := r.Db.Conn.QueryRow(
		`INSERT INTO persons (name,surname,) VALUES($1) RETURNING Id`, per.Name).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed Insert to 'persons' table: ", zap.Error(err))
		return err
	}
	per.Id = lastInsertedId
	return nil
}

func (r *PersonRepository) GetByName(name string) (*domain.Person, error) {
	var result domain.Person
	err := r.Db.Conn.QueryRow(
		`SELECT * FROM persons WHERE name=$1`, name).Scan(&result)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
		return nil, err
	}

	return &result, nil
}

func (r *PersonRepository) GetBySurname(surname string) (*domain.Person, error) {
	var result domain.Person
	err := r.Db.Conn.QueryRow(
		`SELECT * FROM persons WHERE surname=$1`, surname).Scan(&result.Id, &result.Name)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
		return nil, err
	}

	return &result, nil
}

func (r *PersonRepository) GetByAge(age int) (*domain.Person, error) {
	var result domain.Person
	err := r.Db.Conn.QueryRow(
		`SELECT * FROM persons WHERE age=$1`, age).Scan(&result.Id, &result.Name)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
		return nil, err
	}

	return &result, nil
}

func (r *PersonRepository) GetByCountry(country string) (*domain.Person, error) {
	var result domain.Person
	err := r.Db.Conn.QueryRow(
		`SELECT * FROM persons WHERE surname=$1`, country).Scan(&result.Id, &result.Name)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
		return nil, err
	}

	return &result, nil
}

func (r *PersonRepository) Update(per *domain.Person) error {
	_, err := r.Db.Conn.Exec(
		`UPDATE persons SET name=$2 WHERE id=$1`, per.Id, per.Name)
	if err != nil {
		logger.Logger.Error("Failed Update in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) Delete(id int64) error {
	_, err := r.Db.Conn.Exec(
		`DELETE FROM persons WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}
