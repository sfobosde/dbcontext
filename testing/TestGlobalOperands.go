package dbcontext_test

import (
	"fmt"

	"github.com/sfobosde/dbcontext/dbcontext"
)

// Testing all global operands: AND, OR, NOT.
func testGlobalOperands(models *model) {
	testAnd(models)
	testOr(models)
	testNot(models)
}

func testAnd(models *model) {
	user := create(models.Users)
	user.Name = "testAnd"
	models.Users.Save(user)

	fetchedUser, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.And(*fields.ID.Equals(user.ID), *fields.Name.Equals("testAnd"))
	}).First()

	if err != nil {
		panic("testAnd: " + fmt.Sprint(err))
	}

	if fetchedUser == nil {
		panic("testAnd: response is nil")
	}
}

// Coverage OR.
func testOr(models *model) {
	first := create(models.Users)
	first.Name = "testOr First"
	models.Users.Save(first)

	second := create(models.Users)
	second.Name = "testOr Second"
	models.Users.Save(second)

	users, err := models.Users.Search().Where(func(g *dbcontext.Operands, f *userSearch) *dbcontext.GLobalFilter {
		return g.Or(*f.Name.Equals("testOr First"), *f.Name.Equals("testOr Second"))
	}).All()

	if err != nil {
		panic("testOr: " + fmt.Sprint(err))
	}

	fmt.Println(users)

	if len(users) != 2 {
		panic("testOr: Expected 2 objects, actual: " + fmt.Sprint(len(users)))
	}

	if !some(users, func(user user) bool { return user.ID == first.ID }) {
		panic("testOr: user not found (first): " + first.ID)
	}

	if !some(users, func(user user) bool { return user.ID == second.ID }) {
		panic("testOr: user not found (second): " + second.ID)
	}
}

// Coverage Not.
func testNot(models *model) {
	first := create(models.Users)
	first.Name = "testNot"
	models.Users.Save(first)

	second := create(models.Users)
	second.Name = "testOr Second"
	models.Users.Save(second)

	users, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.Not(*operands.Or(*fields.Name.Equals("testNot"), *fields.ID.Equals(second.ID)))
	}).Size(1000).All()

	if err != nil {
		panic("testNot: " + fmt.Sprint(err))
	}

	if some(users, func(value user) bool { return value.ID == first.ID }) {
		panic("testNot: expected no such user (first). actual: found")
	}

	if some(users, func(value user) bool { return value.ID == second.ID }) {
		panic("testNot: expected no such user (second). actual: found")
	}
}
