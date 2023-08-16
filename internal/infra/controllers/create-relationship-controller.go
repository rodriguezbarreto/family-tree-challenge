package controllers

import (
	"encoding/json"
	usecases "family-tree-challenge/internal/use-cases"
	"family-tree-challenge/internal/use-cases/dto"
	"net/http"
)

type CreateRelationshipController struct {
	useCase *usecases.CreateRelationship
}

func NewCreateRelationshipController(useCaseCreateRelationship *usecases.CreateRelationship) *CreateRelationshipController {
	return &CreateRelationshipController{useCase: useCaseCreateRelationship}
}

func (c *CreateRelationshipController) Handler(response http.ResponseWriter, request *http.Request) {
	var input dto.RelationshipInputDTO
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.useCase.Execute(input)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
}
