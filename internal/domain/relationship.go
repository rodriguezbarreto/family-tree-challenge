package domain

import (
	structErrors "family-tree-challenge/internal/domain/errors"

	"github.com/google/uuid"
)

type Relationship struct {
	ID       string `json:"id" validate:"required"`
	Children Person `json:"children"`
	Parent   Person `json:"parent"`
}

func NewRelationship(children Person, parent Person) (*Relationship, error) {
	relationship := &Relationship{
		ID:       uuid.New().String(),
		Children: children,
		Parent:   parent,
	}

	err := structErrors.ValidatorStruct(relationship)
	if err == nil {
		return relationship, nil
	}

	return nil, err
}
