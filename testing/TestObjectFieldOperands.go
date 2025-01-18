package dbcontext_test

import (
	"fmt"

	"github.com/sfobosde/dbcontext/dbcontext"
)

// Coverage ObjectFieldOperands actions.
func testObjectFieldsOperands(models *model) {
	testLink(models)
}

func testLink(models *model) {
	user := create(models.Users)
	user.Name = "testLink"
	models.Users.Save(user)

	group := create(models.Groups)
	group.Name = "testLink group"
	models.Groups.Save(group)

	userGroupMember := create(models.UserGroup)
	userGroupMember.UserId = user.Id
	userGroupMember.GroupId = group.Id
	models.UserGroup.Save(userGroupMember)

	fetchedUserGroup, err := models.UserGroup.Search().Where(func(operands *dbcontext.Operands, fields *userGroupSearch) *dbcontext.GLobalFilter {
		return fields.UserId.Link(user)
	}).Size(1000).All()

	if err != nil {
		panic("testLink: " + fmt.Sprint(err))
	}

	if !some(fetchedUserGroup, func(value userGroup) bool { return value.Id == userGroupMember.Id }) {
		panic("testLink: cannot find required linked object")
	}
}
