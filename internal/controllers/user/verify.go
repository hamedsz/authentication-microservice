package user

import (
	"auth_micro/helpers/auth"
	user "auth_micro/internal/models/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type verifyResponse struct {
	User  interface{}
}

func (controller Controller) Verify(c *gin.Context)  {
	climbs := auth.AuthorizeJWT(c)
	if climbs == nil {
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
		User: result.ToJson(),
	})
}
