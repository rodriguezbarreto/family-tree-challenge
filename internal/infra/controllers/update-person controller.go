package controllers

import (
	"encoding/json"
	usecases "family-tree-challenge/internal/use-cases"
	"family-tree-challenge/internal/use-cases/dto"
	"net/http"

	"github.com/go-chi/chi"
)

type UpdatePersonController struct {
	useCase *usecases.UpdatePerson
}

func NewUpdatePersonController(useCase *usecases.UpdatePerson) *UpdatePersonController {
	return &UpdatePersonController{
		useCase: useCase,
	}
}

func (c *UpdatePersonController) Handler(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	id := chi.URLParam(request, "id")

	var input dto.PersonInputDTO
	err := json.NewDecoder(request.Body).Decode(&input)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		// TODO: AJUSTAR TRATAMENTO DE ERRO
		return
	}

	err = c.useCase.Execute(&id, &input.Name)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		// TODO: AJUSTAR TRATAMENTO DE ERRO
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Atualização bem-sucedida"))
}
