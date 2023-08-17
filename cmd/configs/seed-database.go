package configs

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/infra/database/repositories"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	personRepo := repositories.NewPersonRepository(db)
	relationshipRepo := repositories.NewRelationshipRepository(db)


	peopleNames := []string{"Sonny", "Ann", "Dunny", "Bruce", "Advick", "Martin", "Phoebe", "Anastasia", "Clark", "Jaqueline", "Oprah", "Eric", "Ellen", "Ariel", "Melody"}
	peopleMap := map[string]*domain.Person{}

	for _, name := range peopleNames {
		person, err := domain.NewPerson(name)
		if err != nil {
			return err
		}
		err = personRepo.Create(person)
		if err != nil {
			return err
		}
		peopleMap[name] = person
	}


	relations := []struct {
		Parent string
		Child  string
	}{
		{"Sonny", "Ann"},
		{"Martin", "Phoebe"},
		{"Anastasia", "Phoebe"},
		{"Anastasia", "Clark"},
		{"Oprah", "Eric"},
		{"Ellen", "Eric"},
		{"Ann", "Dunny"},
		{"Ann", "Bruce"},
		{"Advick", "Dunny"},
		{"Advick", "Bruce"},
		{"Phoebe", "Bruce"},
		{"Phoebe", "Dunny"},
		{"Clark", "Jaqueline"},
		{"Eric", "Jaqueline"},
		{"Eric", "Melody"},
		{"Ariel", "Melody"},
	}

	for _, relation := range relations {
		parent := peopleMap[relation.Parent]
		child := peopleMap[relation.Child]
		relationship, err := domain.NewRelationship(*child, *parent)
		if err != nil {
			return err
		}
		err = relationshipRepo.Create(relationship)
		if err != nil {
			return err
		}
	}

	return nil
}
