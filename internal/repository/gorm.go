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

func (r *PersonRepository) GetByName(name string) *domain.Person {
	var person []domain.Person
	r.Db.Gorm.Where("name=?", name).Find(&person)
	return nil
}

func (r *PersonRepository) UpdateName(name string) *domain.Person {
	var person []domain.Person
	r.Db.Gorm.Model(&person).Updates(domain.Person{Name: "John Doe", Age: 35})
	return nil
}
