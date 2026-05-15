package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=18"`
}

func main() {
	// new User
	u := User{
		Name:  "Tips Go",
		Email: "tipsgo@.com",
		Age:   17,
	}
	v := validator.New()
	err := v.Struct(u)

	if err != nil {
		log.Println("Validation failed:")

		for _, e := range err.(validator.ValidationErrors) {
			fmt.Printf("Field: %s, Error: %s \n", e.Field(), e.Tag())
		}
	} else {
		log.Println("Validation successful")
	}
}
