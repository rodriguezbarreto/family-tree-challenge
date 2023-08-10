package usecases_test

import (
	"errors"
	"famlily-tree-challenge/internal/domain"
	usecases "famlily-tree-challenge/internal/use-cases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ListPersons_Success(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	expectedPersons := []*domain.Person{
		{ID: "1", Name: "John"},
		{ID: "2", Name: "Jane"},
	}

	repoMock.On("List", mock.Anything).Return(expectedPersons, nil)

	useCase := usecases.NewListPersons(repoMock)
	persons, err := useCase.Execute(nil)

	assert.NoError(err)
	assert.Equal(expectedPersons, persons)
	repoMock.AssertExpectations(t)
}

func Test_ListPersons_FilterByID(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	filterID := "1"
	expectedPersons := []*domain.Person{
		{ID: "1", Name: "John"},
	}

	repoMock.On("List", &filterID).Return(expectedPersons, nil)

	useCase := usecases.NewListPersons(repoMock)
	persons, err := useCase.Execute(&filterID)

	assert.NoError(err)
	assert.Equal(expectedPersons, persons)
	repoMock.AssertExpectations(t)
}

func Test_ListPersons_RepositoryError(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	expectedErr := errors.New("Failed to list persons")
	repoMock.On("List", mock.Anything).Return(nil, expectedErr)

	useCase := usecases.NewListPersons(repoMock)
	persons, err := useCase.Execute(nil)

	assert.Error(err)
	assert.Nil(persons)
	assert.Equal(err, expectedErr)
	repoMock.AssertExpectations(t)
}
