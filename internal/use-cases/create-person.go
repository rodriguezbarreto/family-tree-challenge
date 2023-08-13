package usecases

import (
	"famlily-tree-challenge/internal/domain"
	"famlily-tree-challenge/internal/use-cases/repositories"
)

type CreatePerson struct {
	respository repositories.PersonRespository
}

func NewCreatePerson(repository repositories.PersonRespository) *CreatePerson {
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
