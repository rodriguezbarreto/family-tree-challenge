package factories

import (
	"family-tree-challenge/internal/infra/controllers"
	"family-tree-challenge/internal/infra/repositories"
	usecases "family-tree-challenge/internal/use-cases"
)

func UpdatePersonFactory() *controllers.UpdatePersonController{
	repository := repositories.NewPersonRepository()
	updatePersonUseCase := usecases.NewUpdatePerson(repository)
	updatePersonController := controllers.NewUpdatePersonController(updatePersonUseCase)

	return updatePersonController
}