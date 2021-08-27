package database

import (
	"auth_micro/helpers/env"
)

var DefaultDatabaseService = env.GetEnv("DB_DEFAULT_SERVICE")

type DatabaseConfig struct {
	Adabter string
	Url     string
	DbName  string
}

var Databases  = map[string]DatabaseConfig{
	"mongodb": {
		Adabter: "mongodb",
		Url:     env.GetEnv("DB_URL"),
		DbName:  env.GetEnv("DB_NAME"),
	},
}

func GetDefault() DatabaseConfig {
	return Databases[DefaultDatabaseService]
}