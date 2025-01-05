package dbcontext_test

import (
	"fmt"
	"time"

	"gitlab.com/dodo141120/dbcontext/dbcontext"
)

// Coverage StringFieldOperands methods.
func testDateTimeFieldOperands(models *model) {
	testDateTimeEquals(models)
	testDateTimeAfter(models)
	testDateTimeBefore(models)
}

// Coverage Equeals filter.
func testDateTimeEquals(models *model) {
	time.Sleep(1 * time.Second)

	user := create(models.Users)
	user.Name = "testDateTimeEquals"
	models.Users.Save(user)

	user, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return fields.Id.Equals(user.Id)
	}).First()

	if err != nil {
		panic("testDateTimeEquals: user by id" + fmt.Sprint(err))
	}

	userByTime, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return fields.CreatedAt.Equals(user.CreatedAt)
	}).First()

	if err != nil {
		panic("testDateTimeEquals user by time" + fmt.Sprint(err))
	}

	if userByTime == nil {
		panic("testDateTimeEquals user by time in nil")
	}

	if userByTime.Id != user.Id {
		panic("testDateTimeEquals users are not equals")
	}
}

// Coverage after filter.
func testDateTimeAfter(models *model) {
	time.Sleep(1 * time.Second)
	startTime := time.Now()
	time.Sleep(1 * time.Second)

	user := create(models.Users)
	user.Name = "testDateTimeAfter"
	models.Users.Save(user)

	userByTime, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return fields.CreatedAt.After(startTime)
	}).First()

	if err != nil {
		panic("testDateTimeAfter: user by id" + fmt.Sprint(err))
	}

	if userByTime.Id != user.Id {
		panic("testDateTimeAfter users are not equals")
	}
}

// Coverage before filter.
func testDateTimeBefore(models *model) {
	time.Sleep(1 * time.Second)
	us := create(models.Users)
	us.Name = "testDateTimeBefore"
	models.Users.Save(us)

	time.Sleep(1 * time.Second)
	startTime := time.Now()

	userByTime, err := models.Users.Search().Where(func(operands *dbcontext.Operands, fields *userSearch) *dbcontext.GLobalFilter {
		return fields.CreatedAt.Before(startTime)
	}).Size(1000).All()

	if err != nil {
		panic("testDateTimeBefore: user by id" + fmt.Sprint(err))
	}

	if !some(userByTime, func(value user) bool { return value.Id == us.Id }) {
		panic("testDateTimeBefore users are not equals (no such user in response)")
	}
}
