package dbcontext_test

import (
	"fmt"

	"github.com/sfobosde/dbcontext/dbcontext"
)

func testLinkedObjectFieldOperands(models *model) {
	testOnCreateWithObjectField(models)
}

// Testing creating and saving objects with additional object field.
func testOnCreateWithObjectField(models *model) {
	userRole := create(models.Roles)
	userRole.Name = "testRole testOnCreateWithObjectField"
	models.Roles.Save(userRole)

	userWithRole := create(models.Users)
	userWithRole.Name = "testOnCreateWithObjectField"

	models.Users.Save(userWithRole)

	userWithRole.Roles = []role{
		*userRole,
	}

	models.Users.Save(userWithRole)

	fetchedRole, err := models.Roles.Fetch(userRole.Id)

	if err != nil {
		panic("testLinkedObjectFieldOperands error in fetchind role")
	}

	if fetchedRole == nil {
		panic("testLinkedObjectFieldOperands role not found")
	}

	fetchedUser, err := models.Users.Fetch(userWithRole.Id)

	if err != nil {
		panic("testLinkedObjectFieldOperands error in fetchind user")
	}

	if fetchedUser == nil {
		panic("testLinkedObjectFieldOperands user not found: " + userWithRole.Id)
	}

	roleFound := false

	for _, role := range fetchedUser.Roles {
		if role.Id == userRole.Id {
			roleFound = true
		}
	}
	if !roleFound {
		panic("testLinkedObjectFieldOperands role not found in user data")
	}

	userByRole, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return fields.Roles.Has(*userRole)
	}).All()

	if err != nil {
		panic("testLinkedObjectFieldOperands error in using has operand")
	}

	if userByRole == nil {
		panic("testLinkedObjectFieldOperands user by role is null")
	}

	if len(userByRole) != 1 {
		panic(fmt.Sprint("testLinkedObjectFieldOperandsexpectd 1 user by role actual: %d", len(userByRole)))
	}

	userByFoleFound := false

	var foundedUserWithRole user

	for _, user := range userByRole {
		if user.Id == userWithRole.Id {
			userByFoleFound = true
			foundedUserWithRole = user
			fmt.Println("testLinkedObjectFieldOperands")
			fmt.Println(user)
		}
	}

	if !userByFoleFound {
		panic("testLinkedObjectFieldOperands not found userByRole")
	}

	roleInFetchedUserFound := false
	for _, role := range foundedUserWithRole.Roles {
		if role.Id == userRole.Id {
			roleInFetchedUserFound = true
		}
	}

	if !roleInFetchedUserFound {
		panic("testLinkedObjectFieldOperands not found role in fetched user")
	}
}
