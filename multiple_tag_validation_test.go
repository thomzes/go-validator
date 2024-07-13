package govalidator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestMultipleTag(t *testing.T) {
	validate := validator.New()

	var user string = "thomas123"

	err := validate.Var(user, "required,number")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTagSuccess(t *testing.T) {
	validate := validator.New()

	var user int = 123

	err := validate.Var(user, "required,number")
	if err != nil {
		fmt.Println(err.Error())
	}
}
