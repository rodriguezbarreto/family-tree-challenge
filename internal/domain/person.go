package domain

import (
	structErrors "family-tree-challenge/internal/domain/errors"

	"github.com/google/uuid"
)

type Person struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"min=3"`
}

func NewPerson(name string) (*Person, error) {

	person := &Person{
		ID:   uuid.New().String(),
		Name: name,
	}

	err := structErrors.ValidatorStruct(person)
	if err == nil {
		return person, nil
	}

	return nil, err
}
