package usecases

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/repositories"
)

type UpdatePerson struct {
	repository repositories.PersonRepository
}

func NewUpdatePerson(repository repositories.PersonRepository) *UpdatePerson {
	return &UpdatePerson{repository: repository}
}

func (u *UpdatePerson) Execute(id *string, name *string) error {
	existingPerson, err := u.repository.List(id)
	if err != nil {
		return err
	}

	if len(existingPerson) == 0 {
		return err
	}

	updatePerson := domain.Person{
		ID:   *id,
		Name: *name,
	}

	err = u.repository.Update(&updatePerson)
	if err != nil {
		return err
	}

	return nil
}
