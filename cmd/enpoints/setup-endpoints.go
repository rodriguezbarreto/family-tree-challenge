package enpoints

import (
	"family-tree-challenge/internal/infra/factories"

	"github.com/go-chi/chi"
)

func SetupEndpoints(router *chi.Mux) {
	// PERSONS
	router.Post("/persons", factories.CreatePersonFactory().Handler)
	router.Get("/persons/{id}", factories.ListPersonFactory().Handler)
	router.Get("/persons", factories.ListPersonFactory().Handler)
	router.Put("/persons/{id}", factories.UpdatePersonFactory().Handler)
	router.Delete("/persons/{id}", factories.DelePersonFactory().Handler)

	// RELATIONSHIPS
	router.Post("/relationships", factories.CreateRelationshipFactory().Handler)
	router.Get("/relationships", factories.ListRelationshipFactory().Handler)
	router.Delete("/relationships/{id}", factories.DeleteRelationshipFactory().Handler)

}
