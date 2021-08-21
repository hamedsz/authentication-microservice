package user

import (
	"auth_micro/helpers/auth"
	"auth_micro/internal/models/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type signupRequest struct{
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,unique_email"`
	Password  string `json:"password" binding:"required"`
}

type signupResponse struct {
	Token string
	User  interface{}
}

func SignUp(c *gin.Context)  {
	var json signupRequest
	err := c.ShouldBind(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(fmt.Sprint("" , json.Password)) , 0)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var model user.User
	model = &user.Model{
		Email:     fmt.Sprint("" , json.Email),
		Password:  string(hashedPassword),
		FirstName: json.FirstName,
		LastName:  json.LastName,
		Role:      "user",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	err = model.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := auth.JWTAuthService().GenerateToken(json.Email , true)

	c.JSON(http.StatusOK , loginResponse{
		Token: token,
		User:  model.ToJson(),
	})
}
