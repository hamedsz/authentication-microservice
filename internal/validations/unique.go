package validations

import (
	"auth_micro/internal/models/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UniqueEmail(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	_ , err := user.Find(gin.H{
		"email": email,
	})

	return err != nil
}
