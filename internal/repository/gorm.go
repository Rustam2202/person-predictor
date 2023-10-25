package repository

import (
	"context"
	"person-predicator/internal/database"
	"person-predicator/internal/domain"
	"person-predicator/internal/logger"

	"go.uber.org/zap"
)

type PersonRepository struct {
	Db          *database.GORM
	personTable string
}

func NewPersonRepository(db *database.GORM) *PersonRepository {
	return &PersonRepository{Db: db, personTable: "persons"}
}

func (r *PersonRepository) Create(ctx context.Context, person *domain.Person) error {
	result := r.Db.Gorm.
		WithContext(ctx).
		Table(r.personTable).
		Create(person)
	if err := result.Error; err != nil {
		logger.Logger.Error("failed to add person to database", zap.Error(err))
		return err
	}
	logger.Logger.Debug("person added to database")
	return nil
}

func (r *PersonRepository) Get(ctx context.Context,
	filters map[string]interface{}, limit int) ([]domain.Person, error) {
	var persons []domain.Person
	if limit <= 0 {
		limit = -1
	}
	result := r.Db.Gorm.
		WithContext(ctx).
		Table(r.personTable).
		Where(filters).
		Limit(limit).
		Find(&persons)
	if err := result.Error; err != nil {
		logger.Logger.Error("failed to get person from database", zap.Error(err))
		return nil, err
	}
	logger.Logger.Debug("person get from database")
	return persons, nil
}

func (r *PersonRepository) Update(ctx context.Context, person *domain.Person) error {
	result := r.Db.Gorm.
		WithContext(ctx).
		Table(r.personTable).
		Model(domain.Person{Id: person.Id}).
		Updates(domain.Person{Name: person.Name})
	if err := result.Error; err != nil {
		logger.Logger.Error("failed to update person from database", zap.Error(err))
		return err
	}
	if result.RowsAffected == 0 {
		logger.Logger.Info("no persons for update")
		return nil
	}
	logger.Logger.Debug("person updated in database")
	return nil
}

func (r *PersonRepository) Delete(ctx context.Context, id int64) error {
	result := r.Db.Gorm.
		WithContext(ctx).
		Table(r.personTable).
		Delete(&domain.Person{Id: id})
	if err := result.Error; err != nil {
		logger.Logger.Error("failed to delete person from database", zap.Error(err))
		return err
	}
	if result.RowsAffected == 0 {
		logger.Logger.Info("no persons for delete")
		return nil
	}
	logger.Logger.Debug("person deleted from database")
	return nil
}
