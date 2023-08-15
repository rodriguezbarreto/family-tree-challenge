package repositories

import "family-tree-challenge/internal/domain"

type RelationshipRespository interface {
	Create(relationship *domain.Relationship) error
	List(relID *string) ([]*domain.Relationship, error)
	Delete(relID string) error
}
