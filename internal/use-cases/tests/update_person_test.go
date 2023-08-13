package usecases_test

import (
	"errors"
	"family-tree-challenge/internal/domain"
	usecases "family-tree-challenge/internal/use-cases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UpdatePerson_Success(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	repoMock.On("List", mock.Anything).Return([]*domain.Person{{ID: "123", Name: "John"}}, nil)
	repoMock.On("Update", mock.Anything).Return(nil)

	useCase := usecases.NewUpdatePerson(repoMock)
	personToUpdate := &domain.Person{ID: "123", Name: "John Updated"}
	err := useCase.Execute(personToUpdate)

	assert.NoError(err)
	repoMock.AssertExpectations(t)
}

func Test_UpdatePerson_RepositoryError(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(PersonRespositoryMock)

	expectedErr := errors.New("Failed to update person")

	repoMock.On("List", mock.Anything).Return([]*domain.Person{{ID: "123", Name: "John"}}, nil)
	repoMock.On("Update", mock.Anything).Return(expectedErr)

	useCase := usecases.NewUpdatePerson(repoMock)
	personToUpdate := &domain.Person{ID: "123", Name: "John Updated"}
	err := useCase.Execute(personToUpdate)

	assert.Error(err)
	assert.Equal(err, expectedErr)
	repoMock.AssertExpectations(t)
}
