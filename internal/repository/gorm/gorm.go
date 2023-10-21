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

func (r *PersonRepository) GetByName(name string) ([]domain.Person, error) {
	var persons []domain.Person
	result := r.Db.Gorm.Where("name=?", name).Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return persons, nil
}

func (r *PersonRepository) GetBySurname(surname string) ([]domain.Person, error) {
	var persons []domain.Person
	result := r.Db.Gorm.Where("surname=?", surname).Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return persons, nil
}

func (r *PersonRepository) GetByAge(age int) ([]domain.Person, error) {
	var persons []domain.Person
	result := r.Db.Gorm.Where("age=?", age).Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return persons, nil
}

func (r *PersonRepository) GetByGender(gender string) ([]domain.Person, error) {
	var persons []domain.Person
	result := r.Db.Gorm.Where("gender=?", gender).Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return persons, nil
}

func (r *PersonRepository) GetByCountry(country string) ([]domain.Person, error) {
	var persons []domain.Person
	result := r.Db.Gorm.Where("country=?", country).Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return persons, nil
}

func (r *PersonRepository) Update(person domain.Person) error {
	r.Db.Gorm.Model(&person).Updates(person)
	return nil
}

func (r *PersonRepository) UpdateName(name string) error {
	var person []domain.Person
	r.Db.Gorm.Model(&person).Updates(domain.Person{Name: "John Doe", Age: 35})
	return nil
}
