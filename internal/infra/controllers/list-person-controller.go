package controllers

import (
	"encoding/json"
	usecases "family-tree-challenge/internal/use-cases"
	"net/http"

	"github.com/go-chi/chi"
)

type ListPersonController struct {
	useCase *usecases.ListPerson
}

func NewListPersonsController(useCaseListPerson *usecases.ListPerson) *ListPersonController {
	return &ListPersonController{
		useCase: useCaseListPerson,
	}
}

func (c *ListPersonController) Handler(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	var idValid *string

	if id != "" {
		idValid = &id
	}

	list, err := c.useCase.Execute(idValid)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(list)
}
