package govalidator

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidationOr(t *testing.T) {
	validate := validator.New()

	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	request := Login{
		Username: "123213",
		Password: "thomas",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok!")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestMustEqualIgnoreCase(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required|numeric"`
		Name     string `validate:"required"`
	}

	request := User{
		Username: "08123123",
		Email:    "thomas@gmail.com",
		Phone:    "08123123",
		Name:     "Gilbert",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validte:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// success
	} else {
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}

}

func TestMusValidRegisterSuccess(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	registerRequest := RegisterRequest{
		Username: "08123123213",
		Email:    "thomas@gmail.com",
		Phone:    "08123123213",
		Password: "Password",
	}

	err := validate.Struct(registerRequest)
	if err != nil {
		fmt.Println(err.Error())
	}

}
