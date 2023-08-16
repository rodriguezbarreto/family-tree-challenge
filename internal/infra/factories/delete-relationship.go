package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func DeleteRelationshipFactory() *controllers.DeleteRelationshipController {
	repository := RelationshipRespositoryFactory()
	deleteRelationshipUseCase := usecases.NewDeleteRelationship(repository)
	deleteRelationshipController := controllers.NewDeleteRelationshipController(deleteRelationshipUseCase)

	return deleteRelationshipController
}
