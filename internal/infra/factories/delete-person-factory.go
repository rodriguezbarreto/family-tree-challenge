package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func DelePersonFactory() *controllers.DeletePersonController {
	repository := PersonRespositoryFactory()
	deletePersonUseCase := usecases.NewDeletePerson(repository)
	deletePersonController := controllers.NewDeletePersonController(deletePersonUseCase)

	return deletePersonController
}
