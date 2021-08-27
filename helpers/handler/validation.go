package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func ValidationErrorHandler(c *gin.Context , err error ) {
	validatorErrs , ok := err.(validator.ValidationErrors)

	var errorMessages []string

	if !ok {
		errorMessages = append(errorMessages , "validation failed")
	}

	for _, e := range validatorErrs {

		errorMessage , ok := errorsMap[e.Tag()]

		if !ok {
			errorMessages = append(errorMessages , err.Error())
			continue
		}

		//error message is ok

		errorMessage = strings.ReplaceAll(
			errorMessage,
			":field",
			e.Field(),
		)

		errorMessages = append(errorMessages , errorMessage)
	}


	c.JSON(http.StatusBadRequest, ErrorResponse{
		Errors: errorMessages,
	})
}

var errorsMap = map[string]string{
	"required": ":field is required.",
	"unique_email": ":field must be unique.",
}