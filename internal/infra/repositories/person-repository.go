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

	person1 := &domain.Person{
		ID:   "mock-id-1",
		Name: "John Doe",
	}

	person2 := &domain.Person{
		ID:   "mock-id-2",
		Name: "Jane Doe",
	}

	if filterByID != nil {
		return []*domain.Person{person1}, nil
	}

	return []*domain.Person{person1, person2}, nil
}

func (r *PersonRespository) Update(person *domain.Person) error {
	// TODO: Implementar conexão com DB
	return nil
}

func (r *PersonRespository) Delete(personID string) error {
	// TODO: Implementar conexão com DB
	return nil
}
