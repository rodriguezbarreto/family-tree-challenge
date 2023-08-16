package usecases

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/dto"
	"family-tree-challenge/internal/use-cases/repositories"
)

type ListRelationship struct {
	repository repositories.RelationshipRespository
}

func NewListRelationship(repository repositories.RelationshipRespository) *ListRelationship {
	return &ListRelationship{repository: repository}
}

func (u *ListRelationship) Execute(filter *dto.RelationshipFilter) ([]*domain.Relationship, error) {
	relationships, err := u.repository.List(filter)
	if err != nil {
		return nil, err
	}

	return relationships, nil
}
