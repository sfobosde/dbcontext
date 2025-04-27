package dbcontext_test

import (
	"github.com/sfobosde/dbcontext/dbcontext"
)

func initModel() *model {
	model := &model{
		Users:     dbcontext.InitDataModel[dbcontext.DataModel[user, userSearch]](),
		UserGroup: dbcontext.InitDataModel[dbcontext.DataModel[userGroup, userGroupSearch]](),
		Groups:    dbcontext.InitDataModel[dbcontext.DataModel[group, groupSearch]](),
		Roles:     dbcontext.InitDataModel[dbcontext.DataModel[role, roleSearch]](),
	}

	return model
}

type model struct {
	Users     *dbcontext.DataModel[user, userSearch]
	UserGroup *dbcontext.DataModel[userGroup, userGroupSearch]
	Groups    *dbcontext.DataModel[group, groupSearch]
	Roles     *dbcontext.DataModel[role, roleSearch]
}

// Структура пользователя.
type user struct {
	dbcontext.BaseEntity

	Name string `json:"name"`
	Male *bool  `json:"male" gorm:"default:null`

	// Roles []role `json:"roles, omitempty" gorm:"many2many:user_roles;joinForeignKey:UserID;joinReferences:RoleID"`
	Roles []role `gorm:"many2many:user_roles;foreignKey:ID;joinForeignKey:UserID;references:ID;joinReferences:RoleID"`
}

type role struct {
	dbcontext.BaseEntity
	Name string `json:"name"`
	// Users []user `json:"users, omitempty" gorm:"many2many:user_roles;joinForeignKey:RoleID;joinReferences:UserID"`
	Users []user `gorm:"many2many:user_roles;foreignKey:ID;joinForeignKey:RoleID;references:ID;joinReferences:UserID"`
}

type roleSearch struct {
	Name *dbcontext.StringFieldOperands `dbcontext:"Name"`
}

type userSearch struct {
	ID        *dbcontext.StringFieldOperands                         `dbcontext:"ID"`
	Login     *dbcontext.StringFieldOperands                         `dbcontext:"Login"`
	Name      *dbcontext.StringFieldOperands                         `dbcontext:"name"`
	CreatedAt *dbcontext.DateTimeFieldOperands                       `dbcontext:"created_at"`
	Male      *dbcontext.BooleanFieldOperands                        `dbcontext:"male"`
	Roles     *dbcontext.LinkedObjectFieldOperands[role, roleSearch] `dbcontext:"Roles:user_roles,user_id,role_id,users.id,roles.id"`
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
