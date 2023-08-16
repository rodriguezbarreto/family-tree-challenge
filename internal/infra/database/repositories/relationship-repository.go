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
	query := r.db.Preload("Child").Preload("Parent")

	if filter.RelID != nil {
		query = query.Where("id = ?", *filter.RelID)
	}
	if filter.ChildID != nil {
		query = query.Where("child_id = ?", *filter.ChildID)
	}
	if filter.ParentID != nil {
		query = query.Where("parent_id = ?", *filter.ParentID)
	}

	tx := query.Find(&list)
	return list, tx.Error
}

func (r *RelationshipRepository) GetByID(relID string) (*domain.Relationship, error) {
	var relationship domain.Relationship

	tx := r.db.Preload("Child").Preload("Parent").First(&relationship, "id = ?", relID)

	return &relationship, tx.Error
}

func (r *RelationshipRepository) Delete(relID string) error {
	tx := r.db.Delete(&domain.Relationship{}, "id = ?", relID)
	return tx.Error
}
