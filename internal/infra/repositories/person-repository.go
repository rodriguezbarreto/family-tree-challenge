package repositories

import (
	"family-tree-challenge/internal/domain"

	"gorm.io/gorm"
)

type PersonRespository struct {
	db *gorm.DB
}

func NewPersonRepository(connetionDB *gorm.DB) *PersonRespository {
	return &PersonRespository{
		db: connetionDB,
	}
}

func (r *PersonRespository) Create(person *domain.Person) error {
	tx := r.db.Create(person)
	return tx.Error
}

func (r *PersonRespository) List(personID *string) ([]*domain.Person, error) {
	var list []*domain.Person

	if personID != nil {
		var person domain.Person
		tx := r.db.First(&person, *personID)
		return []*domain.Person{&person}, tx.Error
	}

	tx := r.db.Find(&list)
	return list, tx.Error
}

func (r *PersonRespository) Update(person *domain.Person) error {
	tx := r.db.Save(person)
	return tx.Error
}

func (r *PersonRespository) Delete(personID string) error {
	tx := r.db.Delete(&domain.Person{}, personID)
	return tx.Error
}
