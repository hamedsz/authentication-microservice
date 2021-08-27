package database

import (
	"auth_micro/config/database"
	"auth_micro/internal/services/database/mongodb"
)

type DatabaseAdabter interface {
	SetConnection(databaseConfig database.DatabaseConfig)
	FindOne(table string , query interface{} , result interface{}) error
	InsertOne(table string , data interface{}) (string , error)
	UpdateOneById(table string , id string ,  data interface{}) error
	DropAll() error
	Count(table string , query interface{}) (int64 , error)
}

func GetAdabter(config database.DatabaseConfig) DatabaseAdabter  {
	switch config.Adabter {
		case "mongodb":
			return mongodb.NewMongoDatabaseAdabter(config , true)
		}

	return nil
}
