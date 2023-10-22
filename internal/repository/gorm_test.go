package repository

import (
	"person-predicator/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockPersonRepository struct {
	mock.Mock
}

func (m *MockPersonRepository) Create(person *domain.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	repo := new(MockPersonRepository)
	person := &domain.Person{
		Name:    "John",
		Surname: "Doe",
	}
	repo.On("Create", person).Return(nil)
	err := repo.Create(person)
	repo.AssertCalled(t, "Create", person)
	assert.NoError(t, err)

	// expectedError := fmt.Errorf("failed to create person")
	repo.On("Create", person).Return(gorm.ErrNotImplemented)
	err = repo.Create(person)
	assert.EqualError(t, err, gorm.ErrNotImplemented.Error())

}
