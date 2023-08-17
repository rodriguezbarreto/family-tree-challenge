package controllers

import (
	"encoding/json"
	usecases "family-tree-challenge/internal/use-cases"
	"family-tree-challenge/internal/use-cases/dto"
	"net/http"
	"strconv"
)

type GetBaconNumberController struct {
	useCase *usecases.GetBaconNumber
}

func NewGetBaconNumberController(useCaseGetBaconNumber *usecases.GetBaconNumber) *GetBaconNumberController {
	return &GetBaconNumberController{
		useCase: useCaseGetBaconNumber,
	}
}

func (c *GetBaconNumberController) Handler(response http.ResponseWriter, request *http.Request) {
	var input dto.BaconNumberInputDTO
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	baconNumber, err := c.useCase.Execute(input.SourceID, input.TargetID)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"baconNumber": strconv.Itoa(baconNumber),
	})
}
