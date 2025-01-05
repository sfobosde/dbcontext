package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/dodo141120/dbcontext/dbcontext"
)

// Coverage where without panics.
func TestWhere(t *testing.T) {
	assert := assert.New(t)

	model := initModel()

	// assert.Equal("Ids", "dfsa")

	assert.NotPanics(func() {
		model.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {

			return operands.And(*fields.Id.Equals("123"))
		})
	})
}

// Coverage create entity.
func TestEntityCreate(t *testing.T) {
	assert := assert.New(t)

	model := initModel()

	assert.NotPanics(func() {
		user := model.Users.Create()

		assert.NotNil(user)
		assert.NotNil(user.Id)
	})
}
