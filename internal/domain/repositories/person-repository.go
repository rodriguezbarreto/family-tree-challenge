package repositories

import "famlily-tree-challenge/internal/domain"

type PersonRespository interface {
	Create(person *domain.Person) error
	List(filterByID *string) ([]*domain.Person, error)
	Update(person *domain.Person) error 
}
