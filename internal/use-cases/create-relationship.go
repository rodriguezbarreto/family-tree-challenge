package usecases

import (
	"errors"
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/dto"
	"family-tree-challenge/internal/use-cases/repositories"
)

type CreateRelationship struct {
	personRepository       repositories.PersonRespository
	relationshipRepository repositories.RelationshipRespository
}

func NewCreateRelationship(personRepo repositories.PersonRespository, relRepo repositories.RelationshipRespository) *CreateRelationship {
	return &CreateRelationship{
		personRepository:       personRepo,
		relationshipRepository: relRepo,
	}
}

func (u *CreateRelationship) Execute(input dto.RelationshipInputDTO) error {
	if input.Child == input.Parent {
		return errors.New("child and parent IDs must be different")
	}

	child, err := u.personRepository.List(&input.Child)
	if err = u.checkPerson(child, err, "child"); err != nil {
		return err
	}

	parent, err := u.personRepository.List(&input.Parent)
	if err = u.checkPerson(parent, err, "parent"); err != nil {
		return err
	}

	newRelationship, err := domain.NewRelationship(*child[0], *parent[0])
	if err != nil {
		return err
	}

	return u.relationshipRepository.Create(newRelationship)
}

func (u *CreateRelationship) checkPerson(person []*domain.Person, err error, role string) error {
	if err != nil {
		return err
	}
	if len(person) == 0 {
		return errors.New(role + " not found")
	}
	return nil
}
