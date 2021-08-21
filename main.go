package main

import (
	"auth_micro/internal/validations"
	"auth_micro/routes"
	"github.com/gin-gonic/gin"
)

func main()  {

	r := gin.Default()
	validations.Register()
	routes.SetRoutes(r)
	r.Run()
	return
}
