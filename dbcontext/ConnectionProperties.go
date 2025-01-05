package dbcontext

import (
	"fmt"
	"os"
)

// Connection Properties params.
type ConnectionProperties struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

// Load connection properties from environment.
// Requires: DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD.
func GetConnectionPropertiesEnv() (*ConnectionProperties, error) {
	conntectionPropertiesMap := map[string]string{
		"DB_HOST":     "",
		"DB_PORT":     "",
		"DB_USER":     "",
		"DB_NAME":     "",
		"DB_PASSWORD": "",
	}

	for key, _ := range conntectionPropertiesMap {
		envValue := os.Getenv(key)
		fmt.Println("Read " + key + ": " + envValue)
		if len(envValue) > 0 {
			conntectionPropertiesMap[key] = envValue
		} else {
			return nil, fmt.Errorf("Fialed to get env value: " + key)
		}
	}

	return &ConnectionProperties{
		Host:     conntectionPropertiesMap["DB_HOST"],
		Port:     conntectionPropertiesMap["DB_PORT"],
		User:     conntectionPropertiesMap["DB_USER"],
		DBName:   conntectionPropertiesMap["DB_NAME"],
		Password: conntectionPropertiesMap["DB_PASSWORD"],
	}, nil
}
