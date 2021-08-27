package user

import (
	"auth_micro/helpers/handler"
	"auth_micro/helpers/handler/errors"
	"auth_micro/internal/models/user"
	"auth_micro/internal/services/auth"
	"auth_micro/internal/validations"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type SignupRequest struct{
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email,unique_email"`
	Password  string `json:"password" binding:"required"`
}

type SignupResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

func (controller Controller) Signup(c *gin.Context)  {
	validations.Register()
	var request SignupRequest
	err := c.BindJSON(&request)

	if err != nil {
		handler.ValidationErrorHandler(c , err)
		return
	}

	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(fmt.Sprint("" , request.Password)) , 0)

	if err != nil {
		handler.ErrorHandler(c , err, errors.INTERNAL_SERVER_ERROR)
		return
	}

	var model user.User
	model = &user.Model{
		Email:     fmt.Sprint("" , request.Email),
		Password:  string(hashedPassword),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Role:      "user",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	err = model.Save()

	if err != nil {
		handler.ErrorHandler(c , err , errors.INTERNAL_SERVER_ERROR)
		return
	}

	token := auth.GetAdabter().GenerateToken(request.Email)

	c.JSON(http.StatusOK , SignupResponse{
		Token: token,
		User:  model,
	})
}
