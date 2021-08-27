package main

import (
	"auth_micro/routes"
	"github.com/gin-gonic/gin"
)

func main()  {

	r := gin.Default()
	routes.SetRoutes(r)
	r.Run()
	return
}
