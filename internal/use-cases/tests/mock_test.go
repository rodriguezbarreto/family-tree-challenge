package usecases_test

import (
	"famlily-tree-challenge/internal/domain"

	"github.com/stretchr/testify/mock"
)

type PersonRespositoryMock struct {
	mock.Mock
}

func (r *PersonRespositoryMock) Create(person *domain.Person) error {
	args := r.Called(person)
	return args.Error(0)
}

func (r *PersonRespositoryMock) List(filterByID *string) ([]*domain.Person, error) {
    args := r.Called(filterByID)
    
    var persons []*domain.Person
    if args.Get(0) != nil {
        persons = args.Get(0).([]*domain.Person)
    }
    
    return persons, args.Error(1)
}