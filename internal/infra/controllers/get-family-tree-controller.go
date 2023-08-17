package controllers

import (
	"encoding/json"
	usecases "family-tree-challenge/internal/use-cases"
	"net/http"

	"github.com/go-chi/chi"
)

type GetGenealogyController struct {
	useCase *usecases.GetGenealogy
}

func NewGetGenealogyController(useCaseGetGenealogy *usecases.GetGenealogy) *GetGenealogyController {
	return &GetGenealogyController{
		useCase: useCaseGetGenealogy,
	}
}

func (c *GetGenealogyController) Handler(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	var idValid *string

	if id != "" {
		idValid = &id
	}
	maxDepht := 3

	// TODO: ESTABELECER REGRA PARA PROFUNDIDADE

	familyTree, err := c.useCase.Execute(idValid, maxDepht)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if familyTree == nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(familyTree)
}
