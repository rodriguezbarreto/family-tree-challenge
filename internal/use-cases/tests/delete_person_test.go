package usecases_test

import (
	"errors"
	"family-tree-challenge/internal/domain"
	usecases "family-tree-challenge/internal/use-cases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_DeletePerson_NotFound(t *testing.T) {
	assert := assert.New(t)
	personRepoMock := new(PersonRespositoryMock)

	personRepoMock.On("List", mock.Anything).Return([]*domain.Person{}, nil)

	useCase := usecases.NewDeletePerson(personRepoMock, nil)
	err := useCase.Execute("123")

	assert.Error(err)
	assert.Equal("person not found", err.Error())
	personRepoMock.AssertExpectations(t)
}

func Test_DeletePerson_RepositoryError(t *testing.T) {
	assert := assert.New(t)
	personRepoMock := new(PersonRespositoryMock)
	relationshipRepoMock := new(RelationshipRepositoryMock)

	expectedErr := errors.New("Failed to delete person")
	personRepoMock.On("List", mock.Anything).Return([]*domain.Person{{ID: "123", Name: "John"}}, nil)
	relationshipRepoMock.On("List", mock.Anything).Return([]*domain.Relationship{}, expectedErr)

	useCase := usecases.NewDeletePerson(personRepoMock, relationshipRepoMock)
	err := useCase.Execute("123")

	assert.Error(err)
	assert.Equal(err, expectedErr)
	personRepoMock.AssertExpectations(t)
	relationshipRepoMock.AssertExpectations(t)
}
