package usecases

import (
	"errors"
	"family-tree-challenge/internal/use-cases/repositories"
)

type DeleteRelationship struct {
	relationshipRepository repositories.RelationshipRespository
}

func NewDeleteRelationship(repository repositories.RelationshipRespository) *DeleteRelationship {
	return &DeleteRelationship{relationshipRepository: repository}
}

func (u *DeleteRelationship) Execute(relID string) error {

	existingRelationship, err := u.relationshipRepository.GetByID(relID)
	if err != nil {
		return err
	}
	
	if existingRelationship == nil {
		return errors.New("relationship not found")
	}

	err = u.relationshipRepository.Delete(relID)
	if err != nil {
		return err
	}

	return nil
}
