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

func (r *PersonRespositoryMock) List() ([]*domain.Person, error) {
	args := r.Called()
	return args.Get(0).([]*domain.Person), args.Error(1)
}
