// Package to read configuration files/params
package config

import (
	"github.com/joho/godotenv"
)

type Environment struct {
	DBHostName string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBPort     string
}

var Env Environment
var ConnectionString string

// ReadConfig reads environment variables from config.env file
func ReadConfig() error {
	if err := godotenv.Load("config.env"); err != nil {
		return err
	}

	return nil
}
