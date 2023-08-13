package usecases

import (
	"errors"
	"family-tree-challenge/internal/use-cases/repositories"
)

type deletePerson struct {
	repository repositories.PersonRespository
}

func NewDeletePerson(repository repositories.PersonRespository) *deletePerson {
	return &deletePerson{repository: repository}
}

func (u *deletePerson) Execute(personID string) error {
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
