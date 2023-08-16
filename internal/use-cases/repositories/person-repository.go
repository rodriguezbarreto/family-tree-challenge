package repositories

import "family-tree-challenge/internal/domain"

type PersonRepository interface {
	Create(person *domain.Person) error
	List(filterByID *string) ([]*domain.Person, error)
	Update(person *domain.Person) error 
	Delete(personID string) error
}
