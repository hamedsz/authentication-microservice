package user

import (
	"auth_micro/helpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id       *string
	FirstName string
	LastName  string
	Role      string
	Email     string
	Password  string

	UpdatedAt interface{}
	CreatedAt interface{}
}

type User interface {
	Save() (error)
	ToJson() (gin.H)
}

func (model *Model) ToJson() (gin.H)  {
	return gin.H{
		"id":         model.Id,
		"email":      model.Email,
		"first_name": model.FirstName,
		"last_name":  model.LastName,
		"role":       model.Role,
		"created_at": model.CreatedAt,
		"updated_at": model.UpdatedAt,
	}
}

func initial(data map[string]interface{}) *Model {
	id            := data["_id"].(primitive.ObjectID).Hex()
	firstName , _ := data["first_name"].(string)
	lastName , _  := data["last_name"].(string)
	email , _     := data["email"].(string)
	password , _  := data["password"].(string)
	role , _      := data["role"].(string)

	return &Model{
		Id: &id,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
		Role: role,
		CreatedAt: data["created_at"],
		UpdatedAt: data["updated_at"],
	}
}

func Find(query interface{}) (*Model, error) {
	var result gin.H
	err := helpers.Database.Collection("users").FindOne(helpers.Ctx , query).Decode(&result)

	if err != nil{
		return nil , err
	}

	return initial(result) , nil
}

func (model *Model) Save()  (error){
	if model.Id == nil{
		err := saveNewUser(model)
		return err
	}else {
		err := updateUser(model)
		return err
	}
}

func saveNewUser(model *Model) (error){
	result , err := helpers.Database.Collection("users").InsertOne(helpers.Ctx , model)
	if err != nil {
		return err
	}

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()

	model.Id = &insertedId

	return nil
}

func updateUser(model *Model) (error){
	objID, err := primitive.ObjectIDFromHex(*model.Id)
	if err != nil {
		return err
	}

	_, err = helpers.Database.Collection("users").UpdateOne(
		helpers.Ctx,
		bson.M{"_id": objID},
		bson.D{
			{
				"$set",
				model,
			},
		},
	)

	return err
}