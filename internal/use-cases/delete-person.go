package usecases

import (
	"errors"
	"family-tree-challenge/internal/use-cases/repositories"
)

type DeletePerson struct {
	repository repositories.PersonRepository
}

func NewDeletePerson(repository repositories.PersonRepository) *DeletePerson {
	return &DeletePerson{repository: repository}
}

func (u *DeletePerson) Execute(personID string) error {
	existingPerson, err := u.repository.List(&personID)
	if err != nil {
		return err
	}

	if len(existingPerson) == 0 {
		return errors.New("person not found")
	}

	err = u.repository.Delete(personID)
	if err != nil {
		return err
	}

	return nil
}
