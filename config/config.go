// Package to read configuration files/params
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	DBHostName       string
	DBUser           string
	DBPassword       string
	DBName           string
	DBConnectionMode string
	DBPort           string
}

var Env Environment
var ConnectionString string

// ReadConfig reads environment variables from config.env file
func ReadConfig() (Environment, error) {
	if err := godotenv.Load("config.env"); err != nil {
		return Env, err
	}

	Env.DBHostName, _ = os.LookupEnv("PGHOST")
	Env.DBUser, _ = os.LookupEnv("PGUSER")
	Env.DBPassword, _ = os.LookupEnv("PGPASSWORD")
	Env.DBPort, _ = os.LookupEnv("PGPORT")
	Env.DBName, _ = os.LookupEnv("PGDATABASE")
	Env.DBConnectionMode, _ = os.LookupEnv("PGSSLMODE")

	// client, err := ent.Open("postgres","host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	ConnectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		Env.DBHostName,
		Env.DBPort,
		Env.DBUser,
		Env.DBName,
		Env.DBPassword,
		Env.DBConnectionMode)

	return Env, nil
}
