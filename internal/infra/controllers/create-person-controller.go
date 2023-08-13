package controllers

import (
	"encoding/json"
	usecases "family-tree-challenge/internal/use-cases"
	"family-tree-challenge/internal/use-cases/dto"
	"net/http"
)

type createPersonController struct {
	useCase *usecases.CreatePerson
}

func NewCreatePersonController(useCaseCreateperson *usecases.CreatePerson) *createPersonController {
	return &createPersonController{useCase: useCaseCreateperson}
}

func (c *createPersonController) Handler(response http.ResponseWriter, request *http.Request) {
	var input dto.CreatePersonInputDTO
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
	}

	err = c.useCase.Execute(input.Name)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusContinue)
}
