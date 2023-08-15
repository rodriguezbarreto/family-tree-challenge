package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"
)

func UpdatePersonFactory() *controllers.UpdatePersonController {
	repository := PersonRespositoryFactory()
	updatePersonUseCase := usecases.NewUpdatePerson(repository)
	updatePersonController := controllers.NewUpdatePersonController(updatePersonUseCase)

	return updatePersonController
}
