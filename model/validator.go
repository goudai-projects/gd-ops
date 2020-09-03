package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(s interface{}) interface{} {
	err := validate.Struct(s)
	errors := err.(validator.ValidationErrors)
	errString := make([]string, len(errors))
	for i, item := range errors {
		errString[i] = fmt.Sprint(item)
	}
	return errString
}

func test() {
	str := Validate(struct {
	}{})

	strings := str.([]string)
	println(strings)

}
