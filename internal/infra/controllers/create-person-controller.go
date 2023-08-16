package controllers

import (
	"encoding/json"
	usecases "family-tree-challenge/internal/use-cases"
	"family-tree-challenge/internal/use-cases/dto"
	"net/http"
)

type CreatePersonController struct {
	useCase *usecases.CreatePerson
}

func NewCreatePersonController(useCaseCreateperson *usecases.CreatePerson) *CreatePersonController {
	return &CreatePersonController{useCase: useCaseCreateperson}
}

func (c *CreatePersonController) Handler(response http.ResponseWriter, request *http.Request) {
	var input dto.PersonInputDTO
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.useCase.Execute(input.Name)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusContinue)
}
