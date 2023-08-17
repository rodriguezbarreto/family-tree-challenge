package controllers

import (
	usecases "family-tree-challenge/internal/use-cases"
	"net/http"

	"github.com/go-chi/chi"
)

type DeleteRelationshipController struct {
	useCase *usecases.DeleteRelationship
}

func NewDeleteRelationshipController(useCase *usecases.DeleteRelationship) *DeleteRelationshipController {
	return &DeleteRelationshipController{
		useCase: useCase,
	}
}

func (c *DeleteRelationshipController) Handler(response http.ResponseWriter, request *http.Request) {
	relID := chi.URLParam(request, "id")

	err := c.useCase.Execute(relID)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		// TODO: AJUSTAR TRATAMENTO DE ERRO
		return
	}

	response.WriteHeader(http.StatusOK)
}
