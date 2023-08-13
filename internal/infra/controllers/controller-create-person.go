package controllers

import (
	"encoding/json"
	usecases "family-tree-challenge/internal/use-cases"
	"family-tree-challenge/internal/use-cases/dto"
	"net/http"
)

type ControllerCreatePerson struct {
	useCase *usecases.CreatePerson
}

func NewControllerCreatePerson(useCaseCreateperson *usecases.CreatePerson) *ControllerCreatePerson {
	return &ControllerCreatePerson{useCase: useCaseCreateperson}
}

func (c *ControllerCreatePerson) Handler(response http.ResponseWriter, request *http.Request) {
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
