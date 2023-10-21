package repository

import (
	"context"
	"person-predicator/internal/database"
	"person-predicator/internal/domain"
	"person-predicator/internal/logger"

	"go.uber.org/zap"
)

type PersonRepository struct {
	// Db *database.Postgres
		Db *database.Postgres

}

func NewPersonRepository(db *database.Postgres) *PersonRepository {
	return &PersonRepository{Db: db}
}

func (r *PersonRepository) Create(ctx context.Context, per *domain.Person) error {
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

func (r *PersonRepository) Get(ctx context.Context, id int64) (*domain.Person, error) {
	var result domain.Person
	err := r.Db.Conn.QueryRow(
		`SELECT * FROM persons WHERE id=$1`, id).Scan(&result)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
		return nil, err
	}

	return &result, nil
}

func (r *PersonRepository) GetByName(ctx context.Context, name string) ([]domain.Person, error) {
	var result []domain.Person
	rows, err := r.Db.Conn.Query(
		`SELECT * FROM persons WHERE name=$1`, name)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))

		return nil, err
	}
	rows.Scan(&result)
	return result, nil
}

func (r *PersonRepository) GetBySurname(ctx context.Context, surname string) ([]domain.Person, error) {
	var result []domain.Person
	err := r.Db.Conn.QueryRow(
		`SELECT * FROM persons WHERE surname=$1`, surname).Scan(&result)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (r *PersonRepository) GetByAge(ctx context.Context, age int) ([]domain.Person, error) {
	var result []domain.Person
	err := r.Db.Conn.QueryRow(
		`SELECT * FROM persons WHERE age=$1`, age).Scan(&result)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (r *PersonRepository) GetByCountry(ctx context.Context, country string) ([]domain.Person, error) {
	var result []domain.Person
	err := r.Db.Conn.QueryRow(
		`SELECT * FROM persons WHERE surname=$1`, country).Scan(&result)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (r *PersonRepository) Update(ctx context.Context, per *domain.Person) error {
	_, err := r.Db.Conn.Exec(
		`UPDATE persons SET name=$2 WHERE id=$1`, per.Id, per.Name)
	if err != nil {
		logger.Logger.Error("Failed Update in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) UpdateName(ctx context.Context, per *domain.Person) error {
	_, err := r.Db.Conn.Exec(
		`UPDATE persons SET name=$2 WHERE id=$1`, per.Id, per.Name)
	if err != nil {
		logger.Logger.Error("Failed Update in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) UpdateSurname(ctx context.Context, per *domain.Person) error {
	_, err := r.Db.Conn.Exec(
		`UPDATE persons SET surname=$2 WHERE id=$1`, per.Id, per.Surname)
	if err != nil {
		logger.Logger.Error("Failed Update in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) UpdateAge(ctx context.Context, per *domain.Person) error {
	_, err := r.Db.Conn.Exec(
		`UPDATE persons SET name=$2 WHERE id=$1`, per.Id, per.Name)
	if err != nil {
		logger.Logger.Error("Failed Update in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) UpdateCountry(ctx context.Context, per *domain.Person) error {
	_, err := r.Db.Conn.Exec(
		`UPDATE persons SET name=$2 WHERE id=$1`, per.Id, per.Name)
	if err != nil {
		logger.Logger.Error("Failed Update in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.Db.Conn.Exec(
		`DELETE FROM persons WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}
