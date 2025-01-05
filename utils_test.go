package main

import (
	"reflect"
	"testing"

	"github.com/sfobosde/dbcontext/dbcontext"
	"github.com/stretchr/testify/assert"
)

// Test connection properties arragning.
func TestGetConnectionConfig(t *testing.T) {
	assert := assert.New(t)

	properties := &dbcontext.ConnectionProperties{
		Host:     "postgres",
		User:     "USER",
		Password: "PASS",
		DBName:   "MYDB",
		Port:     "12345",
	}

	expectDsn := "host=postgres user=USER password=PASS dbname=MYDB port=12345 sslmode=disable TimeZone=UTC"

	assert.EqualValues(expectDsn, dbcontext.GetConnectionConfig(properties), "Test connection properties arragning.")
}

// Coverage get field tag value.
func TestGetFieldTag(test *testing.T) {
	assert := assert.New(test)

	userSearch := userSearch{Id: &dbcontext.StringFieldOperands{}}

	// v := reflect.ValueOf(userSearch)
	t := reflect.TypeOf(userSearch)

	field := t.Field(0)

	tagValue, hasTag := dbcontext.GetFieldTag(field, "dbcontext")
	assert.True(hasTag)
	assert.Equal("Id", tagValue, "Coverage get field tag value.")
}

// Coverage setting field code name from tag value.
func TestSetFieldCodes(test *testing.T) {
	assert := assert.New(test)

	userSearch := &userSearch{Id: &dbcontext.StringFieldOperands{}, Login: &dbcontext.StringFieldOperands{}}

	assert.NotPanics(func() {
		dbcontext.SetFieldCodes(userSearch)
	})
	assert.Equal("Id", userSearch.Id.FieldName, "Check id field filled by metadata.")
	assert.Equal("Login", userSearch.Login.FieldName, "Check login field filled by metadata.")
}

// Coverage set object property.
func TestSetObjectProperty(t *testing.T) {
	assert := assert.New(t)

	user := new(user)

	userId := "IT IS USER ID"

	err := dbcontext.SetObjectProperty(user, "Id", userId)

	assert.Nil(err)
	assert.Equal(userId, user.Id, "Coverage setting object proprties by value.")
}
