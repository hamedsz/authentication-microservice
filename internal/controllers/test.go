package controllers

import (
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context)  {
	var json map[string]interface{}
	c.BindJSON(&json)

	c.JSON(200 , gin.H{
		"status": "ok",
	})
}


