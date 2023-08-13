package usecases_test

import (
	"errors"
	"family-tree-challenge/internal/domain"
	usecases "family-tree-challenge/internal/use-cases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_DeletePerson_Success(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	repoMock.On("List", mock.Anything).Return([]*domain.Person{{ID: "123", Name: "John"}}, nil)
	repoMock.On("Delete", mock.Anything).Return(nil)

	useCase := usecases.NewDeletePerson(repoMock)
	err := useCase.Execute("123")

	assert.NoError(err)
	repoMock.AssertExpectations(t)
}

func Test_DeletePerson_NotFound(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	repoMock.On("List", mock.Anything).Return([]*domain.Person{}, nil)

	useCase := usecases.NewDeletePerson(repoMock)
	err := useCase.Execute("123")

	assert.Error(err)
	assert.Equal("person not found", err.Error())
	repoMock.AssertExpectations(t)
}

func Test_DeletePerson_RepositoryError(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	expectedErr := errors.New("Failed to delete person")
	repoMock.On("List", mock.Anything).Return([]*domain.Person{{ID: "123", Name: "John"}}, nil)
	repoMock.On("Delete", mock.Anything).Return(expectedErr)

	useCase := usecases.NewDeletePerson(repoMock)
	err := useCase.Execute("123")

	assert.Error(err)
	assert.Equal(err, expectedErr)
	repoMock.AssertExpectations(t)
}
