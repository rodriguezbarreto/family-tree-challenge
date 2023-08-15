package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	"family-tree-challenge/internal/infra/repositories"
	usecases "family-tree-challenge/internal/use-cases"
)

func DelePersonFactory() *controllers.DeletePersonController {
	repository := repositories.NewPersonRepository()
	deletePersonUseCase := usecases.NewDeletePerson(repository)
	deletePersonController := controllers.NewDeletePersonController(deletePersonUseCase)

	return deletePersonController
}
