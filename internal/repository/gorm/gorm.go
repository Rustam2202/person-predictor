package repository

import (
	"person-predicator/internal/database"
	"person-predicator/internal/domain"
)

type PersonRepository struct {
	Db *database.GORM
}

func NewPersonRepository(db *database.GORM) *PersonRepository {
	return &PersonRepository{Db: db}
}

func (r *PersonRepository) Create(person *domain.Person) error {
	result := r.Db.Gorm.Create(person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PersonRepository) Get(filters map[string]interface{}, limit int) ([]domain.Person, error) {
	var persons []domain.Person
	result := r.Db.Gorm.Where(filters).Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return persons, nil
}

func (r *PersonRepository) Update(person *domain.Person) error {
	result := r.Db.Gorm.Save(person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PersonRepository) Delete(id int64) error {
	result := r.Db.Gorm.Delete(&domain.Person{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
