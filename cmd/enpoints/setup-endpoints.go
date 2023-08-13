package enpoints

import (
	"family-tree-challenge/internal/infra/controllers"
	"family-tree-challenge/internal/infra/repositories"
	usecases "family-tree-challenge/internal/use-cases"

	"github.com/go-chi/chi"
)

func SetupEndpoints(router *chi.Mux) {

	repository := repositories.NewPersonRepository()
	createPersonUseCase := usecases.NewCreatePerson(repository)
	listPersonUseCase := usecases.NewListPersons(repository)

	controllerCreatePerson := controllers.NewCreatePersonController(createPersonUseCase)
	listPersonController := controllers.NewListPersonsController(listPersonUseCase)


	router.Post("/persons", controllerCreatePerson.Handler)
	router.Get("/persons/{id}", listPersonController.Handler)
	router.Get("/persons", listPersonController.Handler)
	
}


