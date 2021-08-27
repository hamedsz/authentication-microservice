package user

import (
	"auth_micro/helpers/handler"
	"auth_micro/helpers/handler/errors"
	"auth_micro/internal/models/user"
	"auth_micro/internal/services/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

func (controller Controller) Login(c *gin.Context)  {
	var request loginRequest
	err := c.BindJSON(&request)
	if err != nil {
		handler.ValidationErrorHandler(c , err)
		return
	}

	user , err := user.Find(gin.H{
		"email": request.Email,
	})

	if err != nil {
		handler.ErrorHandler(c , err , errors.WRONG_LOGIN_INFO)
		return
	}


	err = bcrypt.CompareHashAndPassword([]byte(user.Password) , []byte(request.Password))

	if err != nil{
		handler.ErrorHandler(c , err , errors.WRONG_LOGIN_INFO)
		return
	}

	token := auth.GetAdabter().GenerateToken(user.Email)

	c.JSON(http.StatusOK , loginResponse{
		Token: token,
		User:  user,
	})
}