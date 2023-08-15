package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	"family-tree-challenge/internal/infra/repositories"
	usecases "family-tree-challenge/internal/use-cases"
)



func ListPersonFactory() *controllers.ListPersonController{
	repository := repositories.NewPersonRepository()
	listPersonUseCase := usecases.NewListPersons(repository)
	listPersonController := controllers.NewListPersonsController(listPersonUseCase)

	return listPersonController
}
