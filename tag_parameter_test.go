package govalidator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestTagParameter(t *testing.T) {
	validate := validator.New()

	user := "10"

	err := validate.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestTagParameterSuccess(t *testing.T) {
	validate := validator.New()

	user := 10

	err := validate.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}

}
