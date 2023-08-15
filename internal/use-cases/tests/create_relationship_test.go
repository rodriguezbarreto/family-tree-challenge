package usecases_test

import (
	"errors"
	"family-tree-challenge/internal/domain"
	usecases "family-tree-challenge/internal/use-cases"
	"family-tree-challenge/internal/use-cases/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateRelationship_Valid(t *testing.T) {
	assert := assert.New(t)

	childID := "child_id"
	parentID := "parent_id"

	child := &domain.Person{ID: childID, Name: "Child"}
	parent := &domain.Person{ID: parentID, Name: "Parent"}

	personRepo := new(PersonRespositoryMock)
	personRepo.On("List", &childID).Return([]*domain.Person{child}, nil)
	personRepo.On("List", &parentID).Return([]*domain.Person{parent}, nil)

	relRepo := new(RelationshipRepositoryMock)
	relRepo.On("Create", mock.Anything).Return(nil)

	useCase := usecases.NewCreateRelationship(personRepo, relRepo)
	input := dto.RelationshipInputDTO{Child: childID, Parent: parentID}

	err := useCase.Execute(input)
	assert.NoError(err)
}

func Test_CreateRelationship_SameIDs(t *testing.T) {
	assert := assert.New(t)

	childID := "child_id"
	parentID := "child_id"

	useCase := usecases.NewCreateRelationship(nil, nil)
	input := dto.RelationshipInputDTO{Child: childID, Parent: parentID}

	err := useCase.Execute(input)
	assert.Equal(errors.New("child and parent IDs must be different"), err)
}

func Test_CreateRelationship_NoChild(t *testing.T) {
	assert := assert.New(t)

	childID := "child_id"
	parentID := "parent_id"

	personRepo := new(PersonRespositoryMock)
	personRepo.On("List", &childID).Return(nil, nil)
	personRepo.On("List", &parentID).Return([]*domain.Person{{ID: parentID, Name: "Parent"}}, nil)

	useCase := usecases.NewCreateRelationship(personRepo, nil)
	input := dto.RelationshipInputDTO{Child: childID, Parent: parentID}

	err := useCase.Execute(input)
	assert.Equal(errors.New("child not found"), err)
}

func Test_CreateRelationship_Execute_NoParent(t *testing.T) {
	assert := assert.New(t)

	childID := "child_id"
	parentID := "parent_id"

	personRepo := new(PersonRespositoryMock)
	personRepo.On("List", &childID).Return([]*domain.Person{{ID: childID, Name: "Child"}}, nil)
	personRepo.On("List", &parentID).Return(nil, nil)

	useCase := usecases.NewCreateRelationship(personRepo, nil)
	input := dto.RelationshipInputDTO{Child: childID, Parent: parentID}

	err := useCase.Execute(input)
	assert.Equal(errors.New("parent not found"), err)
}
