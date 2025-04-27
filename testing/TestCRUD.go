package dbcontext_test

import (
	"fmt"

	"github.com/sfobosde/dbcontext/dbcontext"
)

// Testing CRUD`s.
func testCRUD(models *model) {
	testCreateSearch(models)
	testSearchAll(models)
	testUpdate(models)
}

// Create, Save, Delete.
func testCreateSearch(models *model) {
	user := create(models.Users)
	user.Name = "TESTUSER"
	models.Users.Save(user)

	searchUser, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.And(*fields.Name.Equals("TESTUSER"))
	}).First()

	if err != nil {
		panic("testCreateSearch: " + fmt.Sprint(err))
	}

	if searchUser == nil {
		panic("testCreateSearch: Ошибка при поиске First.")
	}

	if searchUser.ID != user.ID {
		panic("testCreateSearch: Не совпадают id")
	}
}

// Grouped funcs for Search().All() requests.
func testSearchAll(models *model) {
	testSelectDefaultSize(models)
	testCustomSize(models)
	testEmptyAll(models)
}

// Fetching enitities with default request size equals 10.
func testSelectDefaultSize(models *model) {
	i := 0
	for i < 15 {
		i++
		user := create(models.Users)
		user.Name = "testSelectDefaultSize"
		models.Users.Save(user)
	}

	selectedUsers, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.And(*fields.Name.Equals("testSelectDefaultSize"))
	}).All()

	if err != nil {
		panic("testSelectDefaultSize:" + fmt.Sprint(err))
	}

	selectedLen := len(selectedUsers)

	if selectedLen != 10 {
		panic("testSelectDefaultSize:" + "Expected 10 objects only. Actual: " + fmt.Sprint(selectedLen))
	}
}

// On custom size of All().
func testCustomSize(models *model) {
	size := 7

	i := 0
	for i < 15 {
		i++
		user := create(models.Users)
		user.Name = "testCustomSize"
		models.Users.Save(user)
	}

	selectedUsers, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.And(*fields.Name.Equals("testCustomSize"))
	}).Size(size).All()

	if err != nil {
		panic("testCustomSize:" + fmt.Sprint(err))
	}

	selectedLen := len(selectedUsers)

	if selectedLen != size {
		panic("testCustomSize:" + "Expected 7 Objects. Actual: " + fmt.Sprint(selectedLen))
	}
}

// Get empty response of All().
func testEmptyAll(models *model) {
	i := 0
	for i < 15 {
		i++
		user := create(models.Users)
		user.Name = "testEmptyAll"
		models.Users.Save(user)
	}

	selectedUsers, err := models.Users.Search().Size(10).All()

	if err != nil {
		panic("testEmptyAll:" + fmt.Sprint(err))
	}

	selectedLen := len(selectedUsers)

	if selectedLen != 10 {
		panic("testEmptyAll:" + "Expected 10, Actual: " + fmt.Sprint(selectedLen))
	}
}

// Coverage UPDATE.
func testUpdate(models *model) {
	user := create(models.Users)
	user.Name = "testUpdate"
	models.Users.Save(user)
	userId := user.ID

	user.Name = "testUpdateName"
	models.Users.Save(user)

	fetchedUser, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return operands.And(*fields.ID.Equals(userId))
	}).First()

	if err != nil {
		panic(err)
	}

	if fetchedUser == nil {
		panic("testUpdate: Fetch response of First() is nil")
	}

	if fetchedUser.Name != user.Name {
		panic("testUpdate: Name field not match.")
	}
}
