package govalidator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidationMaps(t *testing.T) {
	type School struct {
		Name string `validate:"required"`
	}

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=1"`
		Scholls   map[string]School `validate:"dive,keys,required,min=2,endkeys"`
	}

	validate := validator.New()
	userRequest := User{
		Id:   "1",
		Name: "Gilbert",
		Addresses: []Address{
			{
				City:    "Bacolod",
				Country: "Filipina",
			},
		},
		Hobbies: []string{"Data Analyst"},
		Scholls: map[string]School{
			"Uni": {
				Name: "University of Phillpines",
			},
			"High School": {
				Name: "Test",
			},
			"Test": {
				Name: "Test",
			},
		},
	}

	err := validate.Struct(userRequest)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestBasicValidationMaps(t *testing.T) {
	type School struct {
		Name string `validate:"required"`
	}

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=1"`
		Scholls   map[string]School `validate:"dive,keys,required,min=2,endkeys"`
		Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=0"`
	}

	validate := validator.New()
	userRequest := User{
		Id:   "1",
		Name: "Gilbert",
		Addresses: []Address{
			{
				City:    "Bacolod",
				Country: "Filipina",
			},
		},
		Hobbies: []string{"Data Analyst"},
		Scholls: map[string]School{
			"Uni": {
				Name: "University of Phillpines",
			},
			"High School": {
				Name: "Test",
			},
			"Test": {
				Name: "Test",
			},
		},
		Wallet: map[string]int{
			"BCA": 100,
			"BRI": 1,
		},
	}

	err := validate.Struct(userRequest)
	if err != nil {
		fmt.Println(err.Error())
	}

}
