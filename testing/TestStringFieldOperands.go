package dbcontext_test

import (
	"fmt"

	"github.com/sfobosde/dbcontext/dbcontext"
)

// Coverage StringFieldOperands methods.
func testStringFieldOperands(models *model) {
	testStringEquals(models)
	testIn(models)
	testLike(models)
}

// Equals test.
func testStringEquals(models *model) {
	user := create(models.Users)
	user.Name = "testStringEquals"
	models.Users.Save(user)

	searchUser, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.And(*fields.Name.Equals("testStringEquals"))
	}).First()

	if err != nil {
		panic("testStringEquals" + fmt.Sprint(err))
	}

	if searchUser == nil {
		panic("testStringEquals: Expected value, actual nil")
	}

	if searchUser.ID != user.ID {
		panic("testStringEquals: fetched object id doesnt matches")
	}
}

// Testing "IN" filter.
func testIn(models *model) {
	userIds := make([]string, 0)

	i := 0
	for i < 10 {
		i++
		user := create(models.Users)
		user.Name = "testIn"
		models.Users.Save(user)
		userIds = append(userIds, user.ID)
	}

	fetchedUsers, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.And(*fields.ID.In(userIds))
	}).All()

	if err != nil {
		panic("testIn" + fmt.Sprint(err))
	}

	selectedLen := len(fetchedUsers)

	if selectedLen != 10 {
		panic("testIn:" + "Expecrted 10 objects, Actual: " + fmt.Sprint(selectedLen))
	}

	for _, fetchedUser := range fetchedUsers {
		if !some(userIds, func(value string) bool { return value == fetchedUser.ID }) {
			panic("testIn: ID not found in fetched: " + fetchedUser.ID)
		}
	}

	fetchedIds := mapArr(fetchedUsers, func(value user) string { return value.ID })
	for _, id := range userIds {
		if !some(fetchedIds, func(value string) bool { return value == id }) {
			panic("testIn: id not found in existed: " + id)
		}
	}
}

// Тестирование Equals.
func testLike(models *model) {
	var matchedUser *user
	var anotherUser *user

	matchedUser = create(models.Users)
	matchedUser.Name = "testLike"
	models.Users.Save(matchedUser)

	anotherUser = create(models.Users)
	anotherUser.Name = "testL1ke"
	models.Users.Save(anotherUser)

	users, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.And(*fields.Name.Like("testLi"))
	}).All()

	if err != nil {
		panic("testLike" + fmt.Sprint(err))
	}

	if len(users) != 1 {
		panic("testL1ke: expected 1 object in response, actual: " + fmt.Sprint(len(users)))
	}

	if !some(users, func(value user) bool { return value.ID == matchedUser.ID }) {
		panic("testL1ke: expected value in list, actual: none ")
	}

	if some(users, func(value user) bool { return value.ID == anotherUser.ID }) {
		panic("testL1ke: extra object in response.")
	}
}
