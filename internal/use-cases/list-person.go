package usecases

import (
	"famlily-tree-challenge/internal/domain"
	"famlily-tree-challenge/internal/domain/repositories"
)

type listPersons struct {
	repository repositories.PersonRespository
}

func NewListPersons(repository repositories.PersonRespository) *listPersons {
	return &listPersons{repository: repository}
}

func (u *listPersons) Execute(filterByID *string) ([]*domain.Person, error) {
	persons, err := u.repository.List(filterByID)
	if err != nil {
		return nil, err
	}

	return persons, nil
}
