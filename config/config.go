// Package to read configuration files/params
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	DBHostName       string
	DBUser           string
	DBPassword       string
	DBConnectionMode string
	DBPort           string
}

var Env Environment

func init() {
	var err error

	Env, err = ReadConfig()
	if err != nil {
		panic(err)
	}
}
func ReadConfig() (Environment, error) {
	if err := godotenv.Load("config.env"); err != nil {
		return Env, err
	}

	Env.DBHostName, _ = os.LookupEnv("DB_DNAME")
	Env.DBConnectionMode, _ = os.LookupEnv("POSTGRES_PASSWORD")
	Env.DBPassword, _ = os.LookupEnv("POSTGRES_HOST_AUTH_METHOD")

	return Env, nil
}
