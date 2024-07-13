package govalidator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidationCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()
	userRequest := User{
		Id:   "1",
		Name: "Thomas",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "Manila",
				Country: "Filipina",
			},
		},
	}

	err := validate.Struct(userRequest)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestValidationBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"dive,required,min=1"`
	}

	validate := validator.New()
	userRequest := User{
		Id:   "1",
		Name: "Thomas",
		Addresses: []Address{
			{
				City:    "Bacolod",
				Country: "Filipina",
			},
			{
				City:    "Manila",
				Country: "Filipina",
			},
		},
		Hobbies: []string{"Tanam Pohon"},
	}

	err := validate.Struct(userRequest)
	if err != nil {
		fmt.Println(err.Error())
	}

}
