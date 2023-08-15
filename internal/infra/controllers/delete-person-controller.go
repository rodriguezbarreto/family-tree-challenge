package controllers

import (
	usecases "family-tree-challenge/internal/use-cases"
	"net/http"

	"github.com/go-chi/chi"
)

type deletePersonController struct {
	useCase *usecases.DeletePerson
}

func NewDeletePersonController(useCase *usecases.DeletePerson) *deletePersonController {
	return &deletePersonController{
		useCase: useCase,
	}
}

func (c *deletePersonController) Handler(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	err := c.useCase.Execute(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		// TODO: AJUSTAR TRATAMENTO DE ERRO
		return
	}

	response.WriteHeader(http.StatusOK)
}
