package domain

import (
	structErrors "family-tree-challenge/internal/domain/errors"

	"github.com/google/uuid"
)

type Relationship struct {
	ID       string `gorm:"primary_key" json:"id" validate:"required"`
	ChildID  string `json:"child_id" validate:"required"`
	ParentID string `json:"parent_id" validate:"required"`
	Child    Person `gorm:"foreignKey:ChildID;references:ID"`
	Parent   Person `gorm:"foreignKey:ParentID;references:ID"`
}

func NewRelationship(child Person, parent Person) (*Relationship, error) {
	relationship := &Relationship{
		ID:       uuid.New().String(),
		ChildID:  child.ID,
		ParentID: parent.ID,
		Child:    child,
		Parent:   parent,
	}

	err := structErrors.ValidatorStruct(relationship)
	if err == nil {
		return relationship, nil
	}

	return nil, err
}
