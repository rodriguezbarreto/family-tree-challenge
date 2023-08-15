package tests

import (
	"family-tree-challenge/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewRelationship_ValidRelationship(t *testing.T) {
	assert := assert.New(t)
	childName := "Jonh Jr."
	parentName := "Jonh Sr."

	child, err := domain.NewPerson(childName)
	assert.NoError(err)

	parent, err := domain.NewPerson(parentName)
	assert.NoError(err)

	rel, err := domain.NewRelationship(*child, *parent)
	assert.NoError(err)
	assert.Equal(childName, rel.Child.Name)
	assert.Equal(parentName, rel.Parent.Name)
}

func Test_NewRelationship_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	childName := "Jonh Jr."
	parentName := "Jonh Sr."

	child, err := domain.NewPerson(childName)
	assert.NoError(err)

	parent, err := domain.NewPerson(parentName)
	assert.NoError(err)

	rel, err := domain.NewRelationship(*child, *parent)
	assert.NoError(err)
	assert.NotNil(rel.ID)
}
