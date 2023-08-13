package usecases_test

import (
	"errors"
	usecases "family-tree-challenge/internal/use-cases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func Test_CreatePerson_Success(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	repoMock.On("Create", mock.Anything).Return(nil)

	useCase := usecases.NewCreatePerson(repoMock)
	err := useCase.Execute("John")

	assert.NoError(err)
	repoMock.AssertExpectations(t)
}

func Test_CreatePerson_RepositoryError(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	expectedErr := errors.New("Failed to create person")
	repoMock.On("Create", mock.Anything).Return(expectedErr)

	useCase := usecases.NewCreatePerson(repoMock)
	err := useCase.Execute("John")

	assert.Error(err)
	assert.Equal(err, expectedErr)
	repoMock.AssertExpectations(t)
}
