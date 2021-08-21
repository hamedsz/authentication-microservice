package routes

import (
	"auth_micro/internal/controllers/user"
	"github.com/gin-gonic/gin"
)


func SetRoutes(r *gin.Engine)  {
	api := r.Group("/api")
	apiV1 := api.Group("/v1")

	apiV1.POST("/users/login", user.Login)
	apiV1.POST("/users/signup", user.SignUp)
	apiV1.POST("/users/verify", user.Verify)
}