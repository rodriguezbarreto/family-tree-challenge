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
	return &GetGenealogy{
		repoRel:    repoRel,
		repoPerson: repoPerson,
	}
}

func (u *GetGenealogy) Execute(personID *string, depthLimit int) (*dto.Genealogy, error) {
	genealogy := &dto.Genealogy{Members: []dto.Member{}}

	err := u.initialActions(personID)
	if err != nil {
		return nil, err
	}

	err = u.getAncestors(personID, &genealogy.Members, depthLimit, 0)
	if err != nil {
		return nil, err
	}

	err = u.getChildren(personID, &genealogy.Members, 0)
	if err != nil {
		return nil, err
	}

	err = u.getSiblingsAndUnclesAndAunts(personID, &genealogy.Members)
	if err != nil {
		return nil, err
	}

	return genealogy, nil
}

func (u *GetGenealogy) getAncestors(personID *string, members *[]dto.Member, depthLimit int, currentDepth int) error {
	println(depthLimit, currentDepth)
	if depthLimit != -1 && currentDepth > depthLimit {
		println("acabou")
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

func (u *GetGenealogy) getChildren(personID *string, members *[]dto.Member, currentDepth int) error {
	person := u.getPersonByID(*personID)
	relationships := u.getRelationshipsByParentID(person.ID)
	
	for _, rel := range relationships {
		childMember := dto.Member{Name: rel.Child.Name, Relationships: []dto.Relationship{
			{Name: person.Name, Relationship: "parent"},
		}}
		
		*members = append(*members, childMember)
	}
	return nil
}

func (u *GetGenealogy) getSiblingsAndUnclesAndAunts(personID *string, members *[]dto.Member) error {
	person := u.getPersonByID(*personID)
	relationships := u.getRelationshipsByChildID(person.ID)

	for _, rel := range relationships {
		siblings := u.getRelationshipsByParentID(rel.Parent.ID)

		for _, siblingRel := range siblings {
			if siblingRel.Child.ID != person.ID {
				siblingMember := dto.Member{Name: siblingRel.Child.Name, Relationships: []dto.Relationship{
					{Name: rel.Parent.Name, Relationship: "parent"},
				}}
				*members = append(*members, siblingMember)
			}
		}

		unclesAndAunts := u.getUnclesAndAunts(rel.Parent.ID)
		*members = append(*members, unclesAndAunts...)
	}
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

func (u *GetGenealogy) getUnclesAndAunts(parentID string) []dto.Member {
	var unclesAndAunts []dto.Member
	parentsOfParent := u.getRelationshipsByChildID(parentID)

	for _, parentOfParentRel := range parentsOfParent {
		siblingsOfParent := u.getRelationshipsByParentID(parentOfParentRel.Parent.ID)
		for _, siblingOfParentRel := range siblingsOfParent {
			if siblingOfParentRel.Child.ID != parentID {
				uncleOrAunt := dto.Member{Name: siblingOfParentRel.Child.Name, Relationships: []dto.Relationship{
					{Name: parentOfParentRel.Parent.Name, Relationship: "parent"},
				}}
				unclesAndAunts = append(unclesAndAunts, uncleOrAunt)
			}
		}
	}
	return unclesAndAunts
}

func (u *GetGenealogy) getRelationshipsByParentID(parentID string) []*domain.Relationship {
	var relationships []*domain.Relationship
	for _, rel := range u.allRels {
		if rel.Parent.ID == parentID {
			relationships = append(relationships, rel)
		}
	}
	return relationships
}



func (u *GetGenealogy) initialActions(personID *string) error {

	_, err := u.repoPerson.List(personID)
	if err != nil {
		return err
	}

	allPersons, err := u.repoPerson.List(nil)
	if err != nil {
		return err
	}
	u.allPersons = allPersons

	allRels, err := u.repoRel.List(&dto.RelationshipFilter{})
	if err != nil {
		return err
	}
	u.allRels = allRels

	return nil
}
