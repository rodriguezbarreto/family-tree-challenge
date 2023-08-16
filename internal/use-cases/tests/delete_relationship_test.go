package usecases_test

import (
	"errors"
	"family-tree-challenge/internal/domain"
	usecases "family-tree-challenge/internal/use-cases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DeleteRelationship_Success(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(RelationshipRepositoryMock)

	relID := "123"
	relationship := &domain.Relationship{ID: relID}
	repoMock.On("GetByID", relID).Return(relationship, nil)
	repoMock.On("Delete", relID).Return(nil)

	useCase := usecases.NewDeleteRelationship(repoMock)
	err := useCase.Execute(relID)

	assert.NoError(err)
	repoMock.AssertExpectations(t)
}

func Test_DeleteRelationship_NotFound(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(RelationshipRepositoryMock)

	relID := "123"
	repoMock.On("GetByID", relID).Return(nil, nil)

	useCase := usecases.NewDeleteRelationship(repoMock)
	err := useCase.Execute(relID)

	assert.Error(err)
	assert.Equal("relationship not found", err.Error())
	repoMock.AssertExpectations(t)
}

func Test_DeleteRelationship_RepositoryError(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(RelationshipRepositoryMock)

	relID := "123"
	expectedErr := errors.New("Failed to delete relationship")
	relationship := &domain.Relationship{ID: relID}
	repoMock.On("GetByID", relID).Return(relationship, nil)
	repoMock.On("Delete", relID).Return(expectedErr)

	useCase := usecases.NewDeleteRelationship(repoMock)
	err := useCase.Execute(relID)

	assert.Error(err)
	assert.Equal(expectedErr, err)
	repoMock.AssertExpectations(t)
}
