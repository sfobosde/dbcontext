package main

import (
	"gitlab.com/dodo141120/dbcontext/dbcontext"
)

func initModel() *model {
	model := &model{
		Users:       dbcontext.InitDataModel[dbcontext.DataModel[user, userSearch]](),
		Permissions: dbcontext.InitDataModel[dbcontext.DataModel[permission, permissionSearch]](),
	}

	return model
}

type model struct {
	Users       *dbcontext.DataModel[user, userSearch]
	Permissions *dbcontext.DataModel[permission, permissionSearch]
}

type user struct {
	dbcontext.BaseEntity

	Name string `json:"name"`
}

type userSearch struct {
	Id    *dbcontext.StringFieldOperands `dbcontext:"Id"`
	Login *dbcontext.StringFieldOperands `dbcontext:"Login"`
	Name  *dbcontext.StringFieldOperands `dbcontext:"name"`
}

type permission struct {
	dbcontext.BaseEntity

	Description string `json:"Description"`
}

type permissionSearch struct {
	Id          *dbcontext.StringFieldOperands `dbcontext:"Id"`
	Description *dbcontext.StringFieldOperands `dbcontext:"Description"`
}
