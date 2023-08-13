package repositories

import (
	"family-tree-challenge/internal/domain"
)

type PersonRespository struct{}

func NewPersonRepository() *PersonRespository {
	return &PersonRespository{}
}

func (r *PersonRespository) Create(person *domain.Person) error {
	// TODO: Implementar conexão com DB
	return nil
}

func (r *PersonRespository) List(filterByID *string) ([]*domain.Person, error) {
	// TODO: Implementar conexão com DB
	return []*domain.Person{}, nil
}

func (r *PersonRespository) Update(person *domain.Person) error {
	// TODO: Implementar conexão com DB
	return nil
}

func (r *PersonRespository) Delete(personID string) error {
	// TODO: Implementar conexão com DB
	return nil
}
