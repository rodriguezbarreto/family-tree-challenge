package tests

import (
	"family-tree-challenge/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)



func Test_NewPerson_CreatePerson(t *testing.T) {
	assert := assert.New(t)
	name := "Jonh"

	person, err := domain.NewPerson(name)

	assert.NoError(err)
	assert.Equal(person.Name, name)
}

func Test_NewPerson_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Jonh"

	person, err := domain.NewPerson(name)

	assert.NoError(err)
	assert.NotNil(person.ID)
}

func Test_NewPerson_MustVaidateNameMin(t *testing.T) {
	assert := assert.New(t)
	name := ""

	_, err := domain.NewPerson(name)

	assert.Equal("name is required with min 3", err.Error())
}
