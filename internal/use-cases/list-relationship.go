package usecases

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/dto"
	"family-tree-challenge/internal/use-cases/repositories"
)

type ListRelationship struct {
	repository repositories.RelationshipRepository
}

func NewListRelationship(repository repositories.RelationshipRepository) *ListRelationship {
	return &ListRelationship{repository: repository}
}

func (u *ListRelationship) Execute(filter *dto.RelationshipFilter) ([]*domain.Relationship, error) {
	relationships, err := u.repository.List(filter)
	if err != nil {
		return nil, err
	}

	return relationships, nil
}
