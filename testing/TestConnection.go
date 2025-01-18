package dbcontext_test

import (
	"fmt"

	"github.com/sfobosde/dbcontext/dbcontext"
	"gorm.io/gorm/logger"
)

// Getting environment connection properties values.
func getConnectionPropertiesEnv() *dbcontext.ConnectionProperties {
	if properties, err := dbcontext.GetConnectionPropertiesEnv(); err != nil {
		panic(err.Error())
	} else {
		return properties
	}
}

// Test connection to DB.
func testConnect(properties *dbcontext.ConnectionProperties) {
	defer func() {
		if r := recover(); r != nil {
			panic("testConnect: " + fmt.Sprint(r))
		}
	}()
	dbcontext.Connect(properties, logger.Info)
	fmt.Println("Database connection open")
}

// Test migrations execution.
func testMigrate() {
	defer func() {
		if r := recover(); r != nil {
			panic("testMigrate: " + fmt.Sprint(r))
		}
	}()
	dbcontext.Migrate()
	fmt.Println("Migrations executed")
}
