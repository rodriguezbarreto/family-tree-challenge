package repositories

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/dto"
)

type RelationshipRepository struct{}

// GetByID implements repositories.RelationshipRespository.
func (*RelationshipRepository) GetByID(relID string) (*domain.Relationship, error) {
	panic("unimplemented")
}

func NewRelationshipRepository() *RelationshipRepository {
	return &RelationshipRepository{}
}

func (r *RelationshipRepository) Create(relationship *domain.Relationship) error {
	// TODO: IMPLEMENTAR CONEXÃO COM BANCO DE DADOS
	return nil
}

func (r *RelationshipRepository) List(filter *dto.RelationshipFilter) ([]*domain.Relationship, error) {
	// TODO: IMPLEMENTAR CONEXÃO COM BANCO DE DADOS
	return []*domain.Relationship{}, nil
}

func (r *RelationshipRepository) Delete(relID string) error {
	// TODO: IMPLEMENTAR CONEXÃO COM BANCO DE DADOS
	return nil
}
