package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func CreateRelationshipFactory() *controllers.CreateRelationshipController {
	personRepository := PersonRespositoryFactory()
	relatioshipRepository := RelationshipRespositoryFactory()
	createRelationshipUseCase := usecases.NewCreateRelationship(personRepository, relatioshipRepository)
	controllerCreateRelationship := controllers.NewCreateRelationshipController(createRelationshipUseCase)

	return controllerCreateRelationship
}
