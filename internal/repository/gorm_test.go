package repository

import (
	"context"
	"fmt"
	"person-predicator/internal/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPersonRepository struct {
	mock.Mock
}

func (m *MockPersonRepository) Create(ctx context.Context, person *domain.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	repo := &MockPersonRepository{}
	person := &domain.Person{
		Name:    "John",
		Surname: "Doe",
	}
	repo.On("Create", person).Return(nil)
	err := repo.Create(ctx, person)
	repo.AssertCalled(t, "Create", person)
	assert.NoError(t, err)

	repo.On("Create", person).Return(fmt.Errorf("database error")) // Error case
	err = repo.Create(ctx, person)
	assert.Error(t, err)
	assert.EqualError(t, err, "database error")
}
