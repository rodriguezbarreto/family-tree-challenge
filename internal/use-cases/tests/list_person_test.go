package usecases_test

import (
	"errors"
	"famlily-tree-challenge/internal/domain"
	usecases "famlily-tree-challenge/internal/use-cases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListPersons_Success(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	// Mock data
	expectedPersons := []*domain.Person{
		{ID: "1", Name: "John"},
		{ID: "2", Name: "Jane"},
	}

	repoMock.On("List").Return(expectedPersons, nil)

	useCase := usecases.NewListPersons(repoMock)
	persons, err := useCase.Execute()

	assert.NoError(err)
	assert.Equal(expectedPersons, persons)
	repoMock.AssertExpectations(t)
}

func Test_ListPersons_RepositoryError(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	expectedErr := errors.New("Failed to list persons")
	repoMock.On("List").Return(nil, expectedErr)

	useCase := usecases.NewListPersons(repoMock)
	persons, err := useCase.Execute()

	assert.Error(err)
	assert.Nil(persons)
	assert.Equal(err, expectedErr)
	repoMock.AssertExpectations(t)
}
