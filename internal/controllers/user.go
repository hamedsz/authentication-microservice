package controllers

import (
	user "auth_micro/internal/models/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func StoreUser(c *gin.Context) {
	var json map[string]interface{}
	c.BindJSON(&json)

	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(fmt.Sprint("" , json["password"])) , 0)

	if err != nil {
		log.Fatal(err)
	}

	var model user.User
	model = &user.Model{
		Email:     fmt.Sprint("" , json["email"]),
		Password:  string(hashedPassword),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	err = model.Save()

	if err != nil {
		fmt.Println(err)
		c.Status(500)
		return
	}
}
