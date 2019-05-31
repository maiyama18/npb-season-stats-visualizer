package db

import (
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

func NewGormDB() (*gorm.DB, error) {
	var emptyEnvVars []string
	dbUser := getEnv("DB_USER", "")
	if dbUser == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_USER")
	}
	dbPassword := getEnv("DB_PASSWORD", "")
	if dbPassword == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_PASSWORD")
	}
	dbHost := getEnv("DB_HOST", "")
	if dbHost == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_HOST")
	}
	dbPort := getEnv("DB_PORT", "")
	if dbPort == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_PORT")
	}
	dbSchema := getEnv("DB_SCHEMA", "")
	if dbSchema == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_SCHEMA")
	}
	if len(emptyEnvVars) > 0 {
		return nil, fmt.Errorf("the following environment variables should be set: %s", strings.Join(emptyEnvVars, ", "))
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbSchema)

	return gorm.Open("mysql", connStr)
}

func getEnv(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}
