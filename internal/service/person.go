package service

import (
	"context"
	"person-predicator/internal/domain"
)

type PersonRepository interface {
	Create(*domain.Person) error
	Get(filters map[string]interface{}, limit int) ([]domain.Person, error)
	Update(*domain.Person) error
	Delete(int64) error

	// Get(context.Context, int64) (*domain.Person, error)
	// GetByName(context.Context, string) ([]domain.Person, error)
	// GetBySurname(context.Context, string) ([]domain.Person, error)
	// GetByAge(context.Context, int) ([]domain.Person, error)
	// GetByCountry(context.Context, string) ([]domain.Person, error)
	// Update(context.Context, *domain.Person) error
	// UpdateName(context.Context, *domain.Person) error
	// UpdateSurname(context.Context, *domain.Person) error
	// // UpdatePatronymic(context.Context, *domain.Person) error
	// UpdateAge(context.Context, *domain.Person) error
	// // UpdateGender(context.Context, *domain.Person) error
	// UpdateCountry(context.Context, *domain.Person) error

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
	err := p.repo.Create(&per)
	if err != nil {
		return 0, err
	}
	return per.Id, nil
}

func (p *PersonService) Get(filters map[string]interface{}, limit int) ([]domain.Person, error) {
	result, err := p.repo.Get(filters, limit)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonService) Update(id int64, name, surname, patronymic string, age int, gender, country string) error {
	err := p.repo.Update(&domain.Person{
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
	err := p.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// func (p *PersonService) Get(ctx context.Context, id int64) (*domain.Person, error) {
// 	result, err := p.repo.Get(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (p *PersonService) GetByName(ctx context.Context, name string) ([]domain.Person, error) {
// 	result, err := p.repo.GetByName(ctx, name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (p *PersonService) GetBySurname(ctx context.Context, surname string) ([]domain.Person, error) {
// 	result, err := p.repo.GetBySurname(ctx, surname)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (p *PersonService) GetByAge(ctx context.Context, age int) ([]domain.Person, error) {
// 	result, err := p.repo.GetByAge(ctx, age)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (p *PersonService) GetByCountry(ctx context.Context, country string) ([]domain.Person, error) {
// 	result, err := p.repo.GetByName(ctx, country)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (p *PersonService) Update(ctx context.Context,
// 	id int64, name, surname, patronymic string, age int, gender, country string) error {
// 	err := p.repo.Update(ctx, &domain.Person{
// 		Id:         id,
// 		Name:       name,
// 		Surname:    surname,
// 		Patronymic: patronymic,
// 		Age:        age,
// 		Gender:     gender,
// 		Country:    country,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (p *PersonService) UpdateName(ctx context.Context, id int64, name string) error {
// 	err := p.repo.UpdateName(ctx, &domain.Person{Id: id, Name: name})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (p *PersonService) UpdateSurname(ctx context.Context, id int64, surname string) error {
// 	err := p.repo.UpdateSurname(ctx, &domain.Person{Id: id, Name: surname})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // func (p *PersonService) UpdatePatronymic(ctx context.Context, id int64, patronymic string) error {
// // 	err := p.repo.UpdatePatronymic(ctx, &domain.Person{Id: id, Patronymic: patronymic})
// // 	if err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }

// func (p *PersonService) UpdateAge(ctx context.Context, id int64, age int) error {
// 	err := p.repo.UpdateAge(ctx, &domain.Person{Id: id, Age: age})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // func (p *PersonService) UpdateGender(ctx context.Context, id int64, gender string) error {
// // 	err := p.repo.UpdateGender(ctx, &domain.Person{Id: id, Gender: gender})
// // 	if err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }

// func (p *PersonService) UpdateCountry(ctx context.Context, id int64, country string) error {
// 	err := p.repo.UpdateCountry(ctx, &domain.Person{Id: id, Country: country})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (p *PersonService) Delete(ctx context.Context, id int64) error {
// 	err := p.repo.Delete(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
