package govalidator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidationCrossFields(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,eqfield=Password"`
	}

	validate := validator.New()
	registerUser := RegisterUser{
		Username:        "thomas@gmail.com",
		Password:        "thomas",
		ConfirmPassword: "thomas",
	}

	err := validate.Struct(registerUser)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on tag", fieldError.Tag(), "error", fieldError.Error())
		}
	}

}
