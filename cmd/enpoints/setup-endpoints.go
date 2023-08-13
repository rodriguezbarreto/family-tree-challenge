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
	controllerCreatePerson := controllers.NewCreatePersonController(createPersonUseCase)

	router.Post("/persons", controllerCreatePerson.Handler)
}
