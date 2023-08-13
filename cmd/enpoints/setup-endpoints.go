package enpoints

import (
	"family-tree-challenge/internal/infra/controllers"
	usecases "family-tree-challenge/internal/use-cases"

	"github.com/go-chi/chi"
)

func SetupEndpoints(router *chi.Mux) {

	createPersonUseCase := usecases.NewCreatePerson(repository)
	controllerCreatePerson := controllers.NewControllerCreatePerson(createPersonUseCase)

	router.Post("/persons", controllerCreatePerson.Handler)
}
