package user

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(c *gin.Context)
	Signup(c *gin.Context)
	Verify(c *gin.Context)
}

type Controller struct {}

func NewUserController() UserController{
	return &Controller{}
}