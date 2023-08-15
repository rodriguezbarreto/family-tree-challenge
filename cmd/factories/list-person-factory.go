package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func ListPersonFactory() *controllers.ListPersonController {
	repository := PersonRespositoryFactory()
	listPersonUseCase := usecases.NewListPersons(repository)
	listPersonController := controllers.NewListPersonsController(listPersonUseCase)

	return listPersonController
}
