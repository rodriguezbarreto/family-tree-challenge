package usecases

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/repositories"
)

type CreatePerson struct {
	respository repositories.PersonRepository
}

func NewCreatePerson(repository repositories.PersonRepository) *CreatePerson {
	return &CreatePerson{respository: repository}
}

func (u *CreatePerson) Execute(name string) error {
	newPerson, err := domain.NewPerson(name)
	if err != nil {
		return err
	}

	err = u.respository.Create(newPerson)
	if err != nil {
		return err
	}

	return nil

}
