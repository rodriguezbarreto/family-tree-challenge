package factories

import "family-tree-challenge/internal/infra/database/repositories"

func RelationshipRespositoryFactory() *repositories.RelationshipRepository {

	repository := repositories.NewRelationshipRepository()
	return repository

}
