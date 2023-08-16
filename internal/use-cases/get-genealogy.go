package usecases

import (
	"family-tree-challenge/internal/use-cases/dto"
	"family-tree-challenge/internal/use-cases/repositories"
)

type GetGenealogy struct {
	repoRel repositories.RelationshipRepository
	repoPerson repositories.PersonRepository
}

func NewGetGenealogy(repoRel repositories.RelationshipRepository, repoPerson repositories.PersonRepository ) *GetGenealogy {
	return &GetGenealogy{repoRel: repoRel, repoPerson: repoPerson}
}

func (u *GetGenealogy) Execute(personID *string, depthLimit int) (*Genealogy, error) {
	genealogy := &Genealogy{Members: []Member{}}
	err := u.getAncestors(personID, &genealogy.Members, depthLimit, 0)
	if err != nil {
		return nil, err
	}

	return genealogy, nil
}

func (u *GetGenealogy) getAncestors(personID *string, members *[]Member, depthLimit int, currentDepth int) error {
	if depthLimit != -1 && currentDepth > depthLimit {
		return nil
	}

	person, err := u.repoPerson.List(personID)
	if err != nil {
		return err
	}

	member := Member{Name: person[0].Name, Relationships: []Relationship{}}
	relationships, err := u.repoRel.List(&dto.RelationshipFilter{ChildID: personID})
	if err != nil {
		return err
	}

	for _, rel := range relationships {
		relationship := Relationship{Name: rel.Parent.Name, Relationship: "parent"}
		member.Relationships = append(member.Relationships, relationship)

		err := u.getAncestors(&rel.Parent.ID, members, depthLimit, currentDepth+1)
		if err != nil {
			return err
		}
	}

	*members = append(*members, member)
	return nil
}

type Genealogy struct {
	Members []Member `json:"members"`
}

type Member struct {
	Name          string         `json:"name"`
	Relationships []Relationship `json:"relationships"`
}

type Relationship struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
}
