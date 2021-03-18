package validator

import "github.com/go-playground/validator"

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(input interface{}) []*ValidationError {
	var errors []*ValidationError

	v := validator.New()
	err := v.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
