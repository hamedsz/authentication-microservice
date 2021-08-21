package validations

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Register() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unique_email", UniqueEmail)
	}
}
