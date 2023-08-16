package usecases_test

import (
	"family-tree-challenge/internal/domain"
	usecases "family-tree-challenge/internal/use-cases"
	"family-tree-challenge/internal/use-cases/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ListRelationship_Execute(t *testing.T) {
	assert := assert.New(t)

	mockRelationships := []*domain.Relationship{
		{ID: "rel1", Child: domain.Person{ID: "child1", Name: "Child 1"}, Parent: domain.Person{ID: "parent1", Name: "Parent 1"}},
		{ID: "rel2", Child: domain.Person{ID: "child2", Name: "Child 2"}, Parent: domain.Person{ID: "parent2", Name: "Parent 2"}},
	}

	repoMock := new(RelationshipRepositoryMock)
	repoMock.On("List", mock.Anything).Return(mockRelationships, nil)

	useCase := usecases.NewListRelationship(repoMock)
	filter := &dto.RelationshipFilter{RelID: nil, ChildID: nil, ParentID: nil}
	result, err := useCase.Execute(filter)

	assert.NoError(err)
	assert.Equal(mockRelationships, result)
	repoMock.AssertExpectations(t)
}
