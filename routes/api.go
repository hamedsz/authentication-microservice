package routes

import (
	"auth_micro/internal/controllers/user"
	"github.com/gin-gonic/gin"
)


func SetRoutes(r *gin.Engine)  {
	api := r.Group("/api")

	apiV1 := api.Group("/v1")

	c := user.NewUserController()

	apiV1.POST("/users/login", c.Login)
	apiV1.POST("/users/signup", c.Signup)
	apiV1.POST("/users/verify", c.Verify)
}