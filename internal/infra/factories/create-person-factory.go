package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func CreatePersonFactory() *controllers.CreatePersonController {
	repository := PersonRespositoryFactory()
	createPersonUseCase := usecases.NewCreatePerson(repository)
	controllerCreatePerson := controllers.NewCreatePersonController(createPersonUseCase)

	return controllerCreatePerson
}
