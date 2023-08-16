package repositories

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/dto"

	"gorm.io/gorm"
)

type RelationshipRepository struct {
	db *gorm.DB
}

func NewRelationshipRepository(connetionDB *gorm.DB) *RelationshipRepository {
	return &RelationshipRepository{
		db: connetionDB,
	}
}

func (r *RelationshipRepository) Create(relationship *domain.Relationship) error {
	tx := r.db.Create(relationship)
	return tx.Error
}

func (r *RelationshipRepository) List(filter *dto.RelationshipFilter) ([]*domain.Relationship, error) {
	var list []*domain.Relationship
	tx := r.db.Find(&list)
	return list, tx.Error

	//TODO: IMPLEMENTAR FILTROS
}

func (r *RelationshipRepository) GetByID(relID string) (*domain.Relationship, error) {
	var relationship domain.Relationship

	tx := r.db.First(&relationship, "id = ?", relID)

	return &relationship, tx.Error
}

func (r *RelationshipRepository) Delete(relID string) error {
	// TODO: IMPLEMENTAR CONEXÃO COM BANCO DE DADOS
	panic("unimplemented")
}
