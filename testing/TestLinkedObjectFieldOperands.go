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

	fetchedRole, err := models.Roles.Fetch(userRole.ID)

	if err != nil {
		panic("testLinkedObjectFieldOperands error in fetchind role")
	}

	if fetchedRole == nil {
		panic("testLinkedObjectFieldOperands role not found")
	}

	fetchedUser, err := models.Users.Fetch(userWithRole.ID)

	if err != nil {
		panic("testLinkedObjectFieldOperands error in fetchind user")
	}

	if fetchedUser == nil {
		panic("testLinkedObjectFieldOperands user not found: " + userWithRole.ID)
	}

	roleFound := false

	for _, role := range fetchedUser.Roles {
		if role.ID == userRole.ID {
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

	for _, user := range userByRole {
		if user.ID == userWithRole.ID {
			userByFoleFound = true
			fmt.Println("testLinkedObjectFieldOperands")
			fmt.Println(user)
		}
	}

	if !userByFoleFound {
		panic("testLinkedObjectFieldOperands not found userByRole")
	}
}
