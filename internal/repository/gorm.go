package repository

import (
	"context"
	"person-predicator/internal/database"
	"person-predicator/internal/domain"
)

type PersonRepository struct {
	Db *database.GORM
}

func NewPersonRepository(db *database.GORM) *PersonRepository {
	return &PersonRepository{Db: db}
}

func (r *PersonRepository) Create(ctx context.Context, person *domain.Person) error {
	result := r.Db.Gorm.WithContext(ctx).Create(person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PersonRepository) Get(ctx context.Context,
	filters map[string]interface{}, limit int) ([]domain.Person, error) {
	var persons []domain.Person
	result := r.Db.Gorm.WithContext(ctx).Where(filters).Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return persons, nil
}

func (r *PersonRepository) Update(ctx context.Context, person *domain.Person) error {
	result := r.Db.Gorm.WithContext(ctx).Save(person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PersonRepository) Delete(ctx context.Context, id int64) error {
	result := r.Db.Gorm.WithContext(ctx).Delete(&domain.Person{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
