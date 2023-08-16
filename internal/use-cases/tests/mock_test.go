package usecases_test

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/dto"

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

func (r *PersonRespositoryMock) Update(person *domain.Person) error {
	args := r.Called(person)
	return args.Error(0)
}

func (r *PersonRespositoryMock) Delete(personID string) error {
	args := r.Called(personID)
	return args.Error(0)
}

type RelationshipRepositoryMock struct {
	mock.Mock
}

func (m *RelationshipRepositoryMock) Create(relationship *domain.Relationship) error {
	args := m.Called(relationship)
	return args.Error(0)
}

func (m *RelationshipRepositoryMock) List(filter *dto.RelationshipFilter) ([]*domain.Relationship, error) {
	args := m.Called(filter)
	return args.Get(0).([]*domain.Relationship), args.Error(1)
}

func (m *RelationshipRepositoryMock) GetByID(relID string) (*domain.Relationship, error) {
	args := m.Called(relID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Relationship), args.Error(1)
}

func (m *RelationshipRepositoryMock) Delete(relID string) error {
	args := m.Called(relID)
	return args.Error(0)
}
