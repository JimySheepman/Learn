package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

type User struct {
	FirstName   string `json:"fname" validate:"alpha"`
	LastName    string `json:"lname" validate:"alpha"`
	Age         uint8  `validate:"gte=20,lte=65"`
	Email       string `json:"e-mail" validate:"required,email"`
	JoiningDate string `validate:"datetime"`
}

var validate *validator.Validate

func main() {

	validate = validator.New()

	user := &User{
		FirstName:   "Test25",
		LastName:    "Test",
		Age:         75,
		Email:       "Badger.Smith@",
		JoiningDate: "005-25-10",
	}

	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		fmt.Println("------ List of tag fields with error ---------")

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println("---------------")
		}
		return
	}
}
