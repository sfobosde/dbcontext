package dbcontext_test

import (
	"fmt"

	"gitlab.com/dodo141120/dbcontext/dbcontext"
)

// Coverage BoolFieldOperands methods.
func testBoolFieldOperands(models *model) {
	testBoolEquals(models)
	testBoolEqualsNull(models)
}

// Coverage equals filter.
func testBoolEquals(models *model) {
	testBoolEqualsTrue(models)
	testBoolEqualsFalse(models)
}

// Coverage equals true.
func testBoolEqualsTrue(models *model) {
	userMale := create(models.Users)
	userMale.Name = "testBoolEqualsTrue"
	male := true
	userMale.Male = &male
	models.Users.Save(userMale)

	searchUser, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.And(*fields.Male.Equals(true))
	}).First()

	if err != nil {
		panic("testBoolEqualsTrue" + fmt.Sprint(err))
	}

	if searchUser == nil {
		panic("testBoolEqualsTrue: Expected value, actual nil")
	}

	if searchUser.Id != userMale.Id {
		panic("testBoolEqualsTrue: fetched object id doesnt matches")
	}
}

// Coverage equals false.
func testBoolEqualsFalse(models *model) {
	userMale := create(models.Users)
	userMale.Name = "testBoolEqualsFalse"
	male := false
	userMale.Male = &male
	models.Users.Save(userMale)

	searchUser, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.And(*fields.Male.Equals(false), *fields.Id.Equals(userMale.Id))
	}).First()

	if err != nil {
		panic("testBoolEqualsFalse" + fmt.Sprint(err))
	}

	if searchUser == nil {
		panic("testBoolEqualsFalse: Expected value, actual nil")
	}

	if searchUser.Id != userMale.Id {
		panic("testBoolEqualsFalse: fetched object id doesnt matches. Actual:" + fmt.Sprint(searchUser))
	}
}

// Coverage Equals null.
func testBoolEqualsNull(models *model) {
	userEmpty := create(models.Users)
	userEmpty.Name = "testBoolEqualsNull"
	models.Users.Save(userEmpty)

	userMale := create(models.Users)
	userMale.Name = "testBoolEqualsFalse"
	male := false
	userMale.Male = &male
	models.Users.Save(userMale)

	searchUsers, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		// return operands.And(*fields.Male.Null(false))
		return operands.And(*fields.Male.EqualsOrNull(false))
	}).Size(1000).All()

	if err != nil {
		panic("testBoolEqualsOrNull" + fmt.Sprint(err))
	}

	if searchUsers == nil {
		panic("testBoolEqualsOrNull: Expected value, actual nil. Res: " + fmt.Sprint(searchUsers))
	}

	if !some(searchUsers, func(value user) bool { return value.Id == userEmpty.Id }) {
		panic("testBoolEqualsOrNull: expected value in list, actual: none (userEmpty)")
	}

	if !some(searchUsers, func(value user) bool { return value.Id == userMale.Id }) {
		panic("testBoolEqualsOrNull: expected value in list, actual: none (userMale)")
	}
}
