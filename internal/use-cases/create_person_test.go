package usecases_test

import (
	"errors"
	"famlily-tree-challenge/internal/domain"
	usecases "famlily-tree-challenge/internal/use-cases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(person *domain.Person) error {
	args := r.Called(person)
	return args.Error(0)
}

func Test_CreatePerson_Success(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(repositoryMock)

	repoMock.On("Create", mock.Anything).Return(nil)

	useCase := usecases.NewCreatePerson(repoMock)
	err := useCase.Execute("John")

	assert.NoError(err)
	repoMock.AssertExpectations(t)
}

func Test_CreatePerson_RepositoryError(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(repositoryMock)

	expectedErr := errors.New("Failed to create person")
	repoMock.On("Create", mock.Anything).Return(expectedErr)

	useCase := usecases.NewCreatePerson(repoMock)
	err := useCase.Execute("John")

	assert.Error(err)
	assert.Equal(err, expectedErr)
	repoMock.AssertExpectations(t)
}
