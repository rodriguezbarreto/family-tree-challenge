package usecases

import (
	"errors"
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/dto"
	"family-tree-challenge/internal/use-cases/repositories"
)

type DeletePerson struct {
	personRepository       repositories.PersonRepository
	relationshipRepository repositories.RelationshipRepository
}

func NewDeletePerson(personRepo repositories.PersonRepository, relRepo repositories.RelationshipRepository) *DeletePerson {
	return &DeletePerson{
		personRepository:       personRepo,
		relationshipRepository: relRepo,
	}
}

func (u *DeletePerson) Execute(personID string) error {
	existingPerson, err := u.personRepository.List(&personID)
	if err != nil {
		return err
	}

	if len(existingPerson) == 0 {
		return errors.New("person not found")
	}

	children, err := u.relationshipRepository.List(&dto.RelationshipFilter{ParentID: &personID})
	if err != nil {
		return err
	}

	for _, child := range children {
		grandparents, err := u.relationshipRepository.List(&dto.RelationshipFilter{ChildID: &child.ParentID})
		if err != nil {
			return err
		}

		for _, grandparent := range grandparents {
			newRelationship, err := domain.NewRelationship(child.Child, grandparent.Parent)
			if err != nil {
				return err
			}

			err = u.relationshipRepository.Create(newRelationship)
			if err != nil {
				return err
			}
		}

		err = u.relationshipRepository.Delete(child.ID)
		if err != nil {
			return err
		}
	}

	relationships, err := u.relationshipRepository.List(&dto.RelationshipFilter{ChildID: &personID})
	if err != nil {
		return err
	}
	for _, relationship := range relationships {
		err = u.relationshipRepository.Delete(relationship.ID)
		if err != nil {
			return err
		}
	}

	err = u.personRepository.Delete(personID)
	if err != nil {
		return err
	}

	return nil
}
