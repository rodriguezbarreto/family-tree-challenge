package enpoints

import (
	"family-tree-challenge/cmd/factories"
	"family-tree-challenge/internal/infra/controllers"
	"family-tree-challenge/internal/infra/repositories"
	usecases "family-tree-challenge/internal/use-cases"

	"github.com/go-chi/chi"
)

func SetupEndpoints(router *chi.Mux) {

	repository := repositories.NewPersonRepository()

	
	updatePersonUseCase := usecases.NewUpdatePerson(repository)
	deletePersonUseCase := usecases.NewDeletePerson(repository)

	
	updatePersonController := controllers.NewUpdatePersonController(updatePersonUseCase)
	deletePersonController := controllers.NewDeletePersonController(deletePersonUseCase)

	router.Post("/persons", factories.CreatePersonFactory().Handler)
	router.Get("/persons/{id}", factories.ListPersonFactory().Handler)
	router.Get("/persons", factories.ListPersonFactory().Handler)
	router.Put("/persons/{id}", updatePersonController.Handler)
	router.Delete("/persons/{id}", deletePersonController.Handler)

}
