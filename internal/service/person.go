package service

import (
	"context"
	"person-predicator/internal/domain"
)

type PersonRepository interface {
	Create(context.Context, *domain.Person) error
	Get(ctx context.Context, filters map[string]interface{}, limit int) ([]domain.Person, error)
	Update(context.Context, *domain.Person) error
	// Update(context.Context, map[string]interface{}) error
	Delete(context.Context, int64) error
}

type PersonService struct {
	repo PersonRepository
}

func NewPersonService(r PersonRepository) *PersonService {
	return &PersonService{repo: r}
}

func (p *PersonService) NewPerson(ctx context.Context,
	name, surname, patronymic string, age int, gender, country string) (int64, error) {
	per := domain.Person{
		Name:       name,
		Surname:    surname,
		Patronymic: patronymic,
		Age:        age,
		Gender:     gender,
		Country:    country,
	}
	err := p.repo.Create(ctx, &per)
	if err != nil {
		return 0, err
	}
	return per.Id, nil
}

func (p *PersonService) Get(ctx context.Context,
	filters map[string]interface{}, limit int) ([]domain.Person, error) {
	result, err := p.repo.Get(ctx, filters, limit)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonService) Update(ctx context.Context,
	id int64, name, surname, patronymic string, age int, gender, country string) error {
	err := p.repo.Update(ctx, &domain.Person{
		Id:         id,
		Name:       name,
		Surname:    surname,
		Patronymic: patronymic,
		Age:        age,
		Gender:     gender,
		Country:    country,
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *PersonService) Delete(ctx context.Context, id int64) error {
	err := p.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
