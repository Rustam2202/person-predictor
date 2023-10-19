package repository

import (
	"person-predicator/internal/domain"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestPersonRepository_Create(t *testing.T) {
    // Create a new instance of the mock controller
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Create a mock object for the DB interface
    mockDB := mocks.NewMockDB(ctrl)

    // Create an instance of PersonRepository with the mock DB
    repo := PersonRepository{Db: mockDB}

    // Create a mock result for the Create method
    mockResult := &gorm.DB{
        Error: nil,
    }

    // Set up the expectation for the Create method on the mock DB
    mockDB.EXPECT().Create(gomock.Any()).Return(mockResult).Times(1)

    // Create a test person object
    person := &domain.Person{
        Name: "John Doe",
        Age:  30,
    }

    // Call the Create method on the repository
    err := repo.Create(person)
    if err != nil {
        t.Errorf("Error creating person: %v", err)
    }
}
