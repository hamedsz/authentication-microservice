package handler

import (
	"auth_micro/helpers/env"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct{
	Errors []string
}

type DebugErrorResponse struct{
	Errors []string
	DebugErrorMessage string
}

func ErrorHandler(c *gin.Context , err error , message string)  {
	var response interface{}

	if env.GetEnv("DEBUG") == "true" {
		response = DebugErrorResponse{
			Errors: []string{
				message,
			},
			DebugErrorMessage: err.Error(),
		}
	} else {
		response = ErrorResponse{
			Errors: []string{
				message,
			},
		}
	}
	
	c.JSON(http.StatusBadRequest, response)
}