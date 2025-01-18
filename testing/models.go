package dbcontext_test

import (
	"github.com/sfobosde/dbcontext/dbcontext"
)

func initModel() *model {
	model := &model{
		Users:       dbcontext.InitDataModel[dbcontext.DataModel[user, userSearch]](),
		Permissions: dbcontext.InitDataModel[dbcontext.DataModel[permission, permissionSearch]](),
		UserGroup:   dbcontext.InitDataModel[dbcontext.DataModel[userGroup, userGroupSearch]](),
		Groups:      dbcontext.InitDataModel[dbcontext.DataModel[group, groupSearch]](),
	}

	return model
}

type model struct {
	Users       *dbcontext.DataModel[user, userSearch]
	Permissions *dbcontext.DataModel[permission, permissionSearch]
	UserGroup   *dbcontext.DataModel[userGroup, userGroupSearch]
	Groups      *dbcontext.DataModel[group, groupSearch]
}

type user struct {
	dbcontext.BaseEntity

	Name string `json:"name"`
	Male *bool  `json:"male" gorm:"default:null`
}

type userSearch struct {
	Id        *dbcontext.StringFieldOperands   `dbcontext:"Id"`
	Login     *dbcontext.StringFieldOperands   `dbcontext:"Login"`
	Name      *dbcontext.StringFieldOperands   `dbcontext:"name"`
	CreatedAt *dbcontext.DateTimeFieldOperands `dbcontext:"created_at"`
	Male      *dbcontext.BooleanFieldOperands  `dbcontext:"male"`
}

type permission struct {
	dbcontext.BaseEntity

	Description string `json:"Description"`
}

type permissionSearch struct {
	Id          *dbcontext.StringFieldOperands `dbcontext:"Id"`
	Description *dbcontext.StringFieldOperands `dbcontext:"Description"`
}

type userGroup struct {
	dbcontext.BaseEntity

	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

type userGroupSearch struct {
	dbcontext.BaseEntity

	GroupId *dbcontext.ObjectFieldOperands[group] `dbcontext:"group_id"`
	UserId  *dbcontext.ObjectFieldOperands[user]  `dbcontext:"user_id"`
}

type group struct {
	dbcontext.BaseEntity

	Name string `json:"name"`
}

type groupSearch struct {
	dbcontext.BaseEntity

	Name *dbcontext.StringFieldOperands `dbcontext:"name"`
}
