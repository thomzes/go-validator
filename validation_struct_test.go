package govalidator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "thomas@gmail.com",
		Password: "thomas",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}
	validate := validator.New()
	userRequest := User{
		Id:   "1",
		Name: "Thomas",
		Address: Address{
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}

	err := validate.Struct(userRequest)
	if err != nil {
		fmt.Println(err.Error())
	}

}
