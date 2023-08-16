package controllers

import (
	"encoding/json"
	usecases "family-tree-challenge/internal/use-cases"
	"family-tree-challenge/internal/use-cases/dto"
	"net/http"
)

type ListRelationshipController struct {
	useCase *usecases.ListRelationship
}

func NewListRelationshipController(useCaseListRelationship *usecases.ListRelationship) *ListRelationshipController {
	return &ListRelationshipController{useCase: useCaseListRelationship}
}

func (c *ListRelationshipController) Handler(response http.ResponseWriter, request *http.Request) {
	var filter dto.RelationshipFilter

	queryParams := request.URL.Query()

	if relID := queryParams.Get("relId"); relID != "" {
		filter.RelID = &relID
	}
	if childID := queryParams.Get("childId"); childID != "" {
		filter.ChildID = &childID
	}
	if parentID := queryParams.Get("parentId"); parentID != "" {
		filter.ParentID = &parentID
	}

	relationships, err := c.useCase.Execute(&filter)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(relationships)
}
