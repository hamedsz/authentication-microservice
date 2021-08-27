package user

import (
	user "auth_micro/internal/models/user"
	"auth_micro/internal/services/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type verifyResponse struct {
	User  user.Model `json:"user"`
}

func getToken(c *gin.Context) (*string , error){

	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")

	if len(BEARER_SCHEMA) > len(authHeader){
//		c.AbortWithStatus(http.StatusUnauthorized)
		return nil , errors.New("wrong header!")
	}

	token := authHeader[len(BEARER_SCHEMA):]

	return &token , nil
}

func (controller Controller) Verify(c *gin.Context)  {

	token , err := getToken(c)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	climbs , err := auth.GetAdabter().Authorize(*token)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	result , err := user.Find(gin.H{
		"email": climbs["name"],
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK , verifyResponse{
		User: result,
	})
}
