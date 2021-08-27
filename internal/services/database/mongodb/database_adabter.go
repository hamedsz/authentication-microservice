package mongodb

import (
	"auth_micro/config/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type MongoDatabaseAdabter struct {
	database mongo.Database
	context  context.Context
}

var caches  = map[database.DatabaseConfig]MongoDatabaseAdabter{}

func NewMongoDatabaseAdabter(config database.DatabaseConfig , useConnectionCache bool) *MongoDatabaseAdabter {

	if useConnectionCache {
		if val , ok := caches[config]; ok {
			return &val
		}
	}

	adapter := MongoDatabaseAdabter{}
	adapter.SetConnection(config)
	caches[config] = adapter
	return &adapter
}

func (adapter *MongoDatabaseAdabter) SetConnection(databaseConfig database.DatabaseConfig)  {
	db , ctx := getNewDatabase(databaseConfig)
	adapter.database = *db
	adapter.context  = ctx
}

func getNewDatabase(config database.DatabaseConfig) (*mongo.Database , context.Context) {
	println("new connection")

	client, err := mongo.
		NewClient(
			options.Client().
				ApplyURI(
					config.Url,
				),
		)

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithCancel(context.Background())

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(config.DbName) , ctx
}

func (adabter *MongoDatabaseAdabter) FindOne(table string , query interface{} , result interface{}) error {
	err := adabter.
		database.
		Collection(table).
		FindOne(adabter.context , query).
		Decode(result)

	return err
}

func (adabter *MongoDatabaseAdabter) Count(table string , query interface{}) (int64 , error) {
	return adabter.
		database.
		Collection(table).
		CountDocuments(
			adabter.context,
			query,
		)
}

func (adabter *MongoDatabaseAdabter) InsertOne(table string , data interface{}) (string , error) {
	result , err := adabter.
		database.
		Collection(table).
		InsertOne(adabter.context , data)

	if err != nil{
		return "" , err
	}

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()

	return insertedId , nil
}

func (adabter *MongoDatabaseAdabter) UpdateOneById(table string , id string, data interface{}) error {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_ , err = adabter.
		database.
		Collection(table).
		UpdateOne(
			adabter.context ,
			bson.M{"_id": objID},
			bson.D{
				{
					"$set",
					data,
				},
			},
		)

	return err
}

func (adabter *MongoDatabaseAdabter) DropAll() error{
	return adabter.database.Drop(adabter.context)
}