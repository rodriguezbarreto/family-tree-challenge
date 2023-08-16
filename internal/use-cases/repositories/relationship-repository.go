package repositories

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/dto"
)

type RelationshipRespository interface {
	Create(relationship *domain.Relationship) error
	List(filter *dto.RelationshipFilter) ([]*domain.Relationship, error)
	GetByID(relID string) (*domain.Relationship, error)
	Delete(relID string) error
}
