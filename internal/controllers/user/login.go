package user

import (
	"auth_micro/helpers/auth"
	"auth_micro/internal/models/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token string
	User  interface{}
}

func (controller Controller) Login(c *gin.Context)  {
	var json loginRequest
	err := c.ShouldBind(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user , err := user.Find(gin.H{
		"email": json.Email,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password) , []byte(json.Password))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := auth.JWTAuthService().GenerateToken(user.Email , true)

	c.JSON(http.StatusOK , loginResponse{
		Token: token,
		User:  user.ToJson(),
	})
}