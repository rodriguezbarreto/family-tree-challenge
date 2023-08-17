package usecases

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/use-cases/dto"
	"family-tree-challenge/internal/use-cases/repositories"
)

type GetGenealogy struct {
	repoRel    repositories.RelationshipRepository
	repoPerson repositories.PersonRepository
	allPersons []*domain.Person
	allRels    []*domain.Relationship
}

func NewGetGenealogy(repoRel repositories.RelationshipRepository, repoPerson repositories.PersonRepository) *GetGenealogy {
	allPersons, _ := repoPerson.List(nil)
	allRels, _ := repoRel.List(&dto.RelationshipFilter{})

	return &GetGenealogy{
		repoRel:    repoRel,
		repoPerson: repoPerson,
		allPersons: allPersons,
		allRels:    allRels,
	}
}

func (u *GetGenealogy) Execute(personID *string, depthLimit int) (*dto.Genealogy, error) {
	genealogy := &dto.Genealogy{Members: []dto.Member{}}
	err := u.getAncestors(personID, &genealogy.Members, depthLimit, 0)
	if err != nil {
		return nil, err
	}

	return genealogy, nil
}

func (u *GetGenealogy) getAncestors(personID *string, members *[]dto.Member, depthLimit int, currentDepth int) error {
	if depthLimit != -1 && currentDepth > depthLimit {
		return nil
	}

	person := u.getPersonByID(*personID)
	member := dto.Member{Name: person.Name, Relationships: []dto.Relationship{}}
	relationships := u.getRelationshipsByChildID(*personID)

	for _, rel := range relationships {
		relationship := dto.Relationship{Name: rel.Parent.Name, Relationship: "parent"}
		member.Relationships = append(member.Relationships, relationship)

		err := u.getAncestors(&rel.Parent.ID, members, depthLimit, currentDepth+1)
		if err != nil {
			return err
		}
	}

	*members = append([]dto.Member{member}, *members...)
	return nil
}

func (u *GetGenealogy) getPersonByID(personID string) *domain.Person {
	for _, person := range u.allPersons {
		if person.ID == personID {
			return person
		}
	}
	return nil
}

func (u *GetGenealogy) getRelationshipsByChildID(childID string) []*domain.Relationship {
	var relationships []*domain.Relationship
	for _, rel := range u.allRels {
		if rel.Child.ID == childID {
			relationships = append(relationships, rel)
		}
	}
	return relationships
}

