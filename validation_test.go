package govalidator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil!")
	}
}

func TestValidationFailed(t *testing.T) {
	validate := validator.New()
	var user string = ""

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestValidationSuccess(t *testing.T) {
	validate := validator.New()
	var user string = "Thomas"

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}

}
