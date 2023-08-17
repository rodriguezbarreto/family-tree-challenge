package controllers

import (
	"encoding/json"
	usecases "family-tree-challenge/internal/use-cases"
	"log"
	"net/http"
	"os"
	"strconv"

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

	maxDepthStr := os.Getenv("MAX_DEPTH")
	maxDepth, err := strconv.Atoi(maxDepthStr)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	familyTree, err := c.useCase.Execute(idValid, maxDepth)
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
