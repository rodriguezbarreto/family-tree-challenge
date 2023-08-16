package usecases

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/repositories"
)

type ListPerson struct {
	repository repositories.PersonRepository
}

func NewListPersons(repository repositories.PersonRepository) *ListPerson {
	return &ListPerson{repository: repository}
}

func (u *ListPerson) Execute(personID *string) ([]*domain.Person, error) {
	persons, err := u.repository.List(personID)
	if err != nil {
		return nil, err
	}

	return persons, nil
}
