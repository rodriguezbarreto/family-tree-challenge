package usecases

import (
	"famlily-tree-challenge/internal/domain"
	"famlily-tree-challenge/internal/domain/repositories"
)

type createPerson struct {
	respository repositories.PersonRespository
}

func NewCreatePerson(repository repositories.PersonRespository) *createPerson {
	return &createPerson{respository: repository}
}

func (u *createPerson) Execute(name string) error {
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
