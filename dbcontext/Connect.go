package dbcontext

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Open DB connection
func Connect(properties *ConnectionProperties, logLevel logger.LogLevel) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			// LogLevel:      logger.Silent, // Log level
			LogLevel:                  logLevel, // Log level
			IgnoreRecordNotFoundError: true,     // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,    // Disable color
		},
	)

	dsn := GetConnectionConfig(properties)
	contextModel := getContextModel()

	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger}); err == nil {
		contextModel.db = db
	} else {
		panic("dbontext: Failed to connect to DB.")
	}
}
