package factories

import (
	"family-tree-challenge/internal/infra/database"
	"family-tree-challenge/internal/infra/database/repositories"
)

func RelationshipRespositoryFactory() *repositories.RelationshipRepository {
	connection := database.PostgresConnection()
	repository := repositories.NewRelationshipRepository(connection)

	return repository

}
