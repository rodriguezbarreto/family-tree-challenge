package configs

import (
	"family-tree-challenge/internal/domain"
	"family-tree-challenge/internal/infra/database/repositories"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	personRepo := repositories.NewPersonRepository(db)
	relationshipRepo := repositories.NewRelationshipRepository(db)


	peopleNames := []string{"Sonny", "Ann", "Dunny", "Bruce", "Advick", "Marting", "Phoebe", "Anastasia", "Clark", "Jaqueline", "Oprah", "Eric", "Ellen", "Ariel", "Melody"}
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
		{"Ann", "Dunny"},
		{"Ann", "Bruce"},
		{"Advick", "Dunny"},
		{"Advick", "Bruce"},
		{"Marting", "Phoebe"},
		{"Anastasia", "Phoebe"},
		{"Phoebe", "Bruce"},
		{"Anastasia", "Clark"},
		{"Clark", "Jaqueline"},
		{"Oprah", "Eric"},
		{"Eric", "Jaqueline"},
		{"Ellen", "Eric"},
		{"Ariel", "Melody"},
		{"Eric", "Melody"},
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
