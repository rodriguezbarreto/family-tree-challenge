package usecases

import (
	"family-tree-challenge/internal/use-cases/dto"
	"family-tree-challenge/internal/use-cases/repositories"
)

type GetFamilyTree struct {
	relationshipRepo repositories.RelationshipRepository
}

func NewGetFamilyTree(repository repositories.RelationshipRepository) *GetFamilyTree {
	return &GetFamilyTree{relationshipRepo: repository}
}

func (u *GetFamilyTree) Execute(personID string, maxDepth int) ([]dto.FamilyMember, error) {
	visited := make(map[string]bool)
	return u.getAncestors(personID, maxDepth, visited)
}

func (u *GetFamilyTree) getAncestors(personID string, depth int, visited map[string]bool) ([]dto.FamilyMember, error) {
	if depth == 0 || visited[personID] {
		return nil, nil
	}

	visited[personID] = true

	rels, err := u.relationshipRepo.List(&dto.RelationshipFilter{ChildID: &personID})
	if err != nil {
		return nil, err
	}

	var familyTree []dto.FamilyMember

	for _, rel := range rels {
		parent := dto.FamilyMember{
			Name: rel.Parent.Name,
			Relationships: []dto.FamilyRelation{
				{
					Name:         rel.Child.Name,
					Relationship: "child",
				},
			},
		}

		ancestors, err := u.getAncestors(rel.Parent.ID, depth-1, visited)
		if err != nil {
			return nil, err
		}

		familyTree = append(familyTree, parent)
		familyTree = append(familyTree, ancestors...)
	}

	return familyTree, nil
}
