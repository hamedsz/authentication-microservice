package user

import (
	config "auth_micro/config/database"
	database "auth_micro/internal/services/database"
	"time"
)

type Model struct {
	Id        string    `bson:"_id,omitempty"        json:"id"`
	FirstName string    `bson:"first_name" json:"first_name"`
	LastName  string    `bson:"last_name"  json:"last_name"`
	Role      string    `bson:"role"       json:"role"`
	Email     string    `bson:"email"      json:"email"`
	Password  string    `bson:"password"   json:"-"`

	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type User interface {
	Save() (error)
}

func Find(query interface{}) (Model, error) {
	var result Model
	err := database.GetAdabter(config.GetDefault()).FindOne("users" , query , &result)
	return result , err
}

func (model *Model) Save()  (error){
	if model.Id == ""{
		err := saveNewUser(model)
		return err
	}else {
		err := updateUser(model)
		return err
	}
}

func saveNewUser(model *Model) (error){
	result , err := database.GetAdabter(config.GetDefault()).InsertOne("users" , model)
	if err != nil {
		return err
	}

	model.Id = result

	return nil
}

func updateUser(model *Model) (error){
	err := database.
		GetAdabter(config.GetDefault()).
		UpdateOneById("users" , model.Id , model)
	return err
}