package repositories

import "famlily-tree-challenge/internal/domain"

type PersonRespository interface {
	Create(person *domain.Person) error
}
