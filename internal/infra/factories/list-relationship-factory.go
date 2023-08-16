package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func ListRelationshipFactory() *controllers.ListRelationshipController {
	repository := RelationshipRespositoryFactory()
	listRelationshipUseCase := usecases.NewListRelationship(repository)
	listRelationshipController := controllers.NewListRelationshipController(listRelationshipUseCase)

	return listRelationshipController
}
