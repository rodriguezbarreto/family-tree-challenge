package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	"family-tree-challenge/internal/infra/repositories"
	usecases "family-tree-challenge/internal/use-cases"
)

func CreatePersonFactory() *controllers.CreatePersonController{
	repository := repositories.NewPersonRepository()
	createPersonUseCase := usecases.NewCreatePerson(repository)
	controllerCreatePerson := controllers.NewCreatePersonController(createPersonUseCase)

	return controllerCreatePerson
}
