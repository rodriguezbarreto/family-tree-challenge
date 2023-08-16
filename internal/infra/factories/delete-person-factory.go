package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func DelePersonFactory() *controllers.DeletePersonController {
	repository := PersonRespositoryFactory()
	repositoryRel := RelationshipRespositoryFactory()
	deletePersonUseCase := usecases.NewDeletePerson(repository, repositoryRel)
	deletePersonController := controllers.NewDeletePersonController(deletePersonUseCase)

	return deletePersonController
}
