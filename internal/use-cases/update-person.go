package usecases

import (
	"errors"
	"famlily-tree-challenge/internal/domain"
	"famlily-tree-challenge/internal/domain/repositories"
)

type updatePerson struct {
	repository repositories.PersonRespository
}

func NewUpdatePerson(repository repositories.PersonRespository) *updatePerson {
	return &updatePerson{repository: repository}
}

func (u *updatePerson) Execute(personToUpdate *domain.Person) error {
	existingPerson, err := u.repository.List(&personToUpdate.ID)
	if err != nil {
		return err
	}

	if len(existingPerson) == 0 {
		return errors.New("person not found")
	}

	err = u.repository.Update(personToUpdate)
	if err != nil {
		return err
	}

	return nil
}
