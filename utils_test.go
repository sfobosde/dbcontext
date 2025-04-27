package main

import (
	"fmt"
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

	userSearch := userSearch{ID: &dbcontext.StringFieldOperands{}}

	// v := reflect.ValueOf(userSearch)
	t := reflect.TypeOf(userSearch)

	field := t.Field(0)

	tagValue, hasTag := dbcontext.GetFieldTag(field, "dbcontext")
	assert.True(hasTag)
	assert.Equal("ID", tagValue, "Coverage get field tag value.")
}

// Coverage setting field code name from tag value.
func TestSetFieldCodes(test *testing.T) {
	assert := assert.New(test)

	userSearch := &userSearch{ID: &dbcontext.StringFieldOperands{}, Login: &dbcontext.StringFieldOperands{}}

	assert.NotPanics(func() {
		dbcontext.SetFieldCodes(userSearch)
	})
	assert.Equal("ID", userSearch.ID.FieldName, "Check id field filled by metadata.")
	assert.Equal("Login", userSearch.Login.FieldName, "Check login field filled by metadata.")
}

// Coverage set object property.
func TestSetObjectProperty(t *testing.T) {
	assert := assert.New(t)

	user := new(user)

	userId := "IT IS USER ID"

	err := dbcontext.SetObjectProperty(user, "ID", userId)

	assert.Nil(err)
	assert.Equal(userId, user.ID, "Coverage setting object proprties by value.")
}

// Coverage get object property value my name.
func TestGetFieldValue(t *testing.T) {
	assert := assert.New(t)

	assert.NotPanics(func() {
		user := new(user)
		user.ID = "r435t345t34"

		userId, err := dbcontext.GetFieldValue(*user, "ID")

		assert.Nil(err, "Check exsited field")
		assert.Equal(userId, "r435t345t34")
	})
}

// Coverage get object property value my name.
func TestGetFieldValueIterable(t *testing.T) {
	assert := assert.New(t)

	assert.NotPanics(func() {
		firstUser := new(user)
		firstUser.ID = "r435t345t34"

		secondUser := new(user)
		secondUser.ID = "r32rrgergg"

		userIds := []string{firstUser.ID, secondUser.ID}

		ids := arrFunc(secondUser, firstUser)

		assert.NotNil(ids)

		for _, value := range ids {
			assert.NotNil(value)
		}

		for _, value := range userIds {
			assert.True(some(ids, func(val any) bool { return value == val }))
		}

		for _, value := range ids {
			assert.True(some(userIds, func(val string) bool { return value == val }))
		}
	})
}

func arrFunc[T any](values ...*T) []any {
	var ids []any
	for _, value := range values {
		id, err := dbcontext.GetFieldValue(*value, "ID")
		fmt.Println(err)
		fmt.Println(value, id)
		ids = append(ids, id)
	}

	return ids
}
func indexOf[T any](arr []T, compare func(value T) bool) int {
	for i, value := range arr {
		if compare(value) {
			return i
		}
	}

	return -1
}

func some[T any](arr []T, compare func(value T) bool) bool {
	return indexOf(arr, compare) != -1
}
