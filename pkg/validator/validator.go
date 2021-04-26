package validator

import (
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type ValidationError struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func ValidateStruct(input interface{}) []*ValidationError {
	v := validator.New()
	en := en.New()
	uniTranslator := ut.New(en, en)

	translator, _ := uniTranslator.GetTranslator("en")

	enTranslations.RegisterDefaultTranslations(v, translator)

	if err := v.Struct(input); err != nil {
		return buildTranslatedErrorMessages(err.(validator.ValidationErrors), translator)
	}

	return nil
}

func buildTranslatedErrorMessages(err validator.ValidationErrors, translator ut.Translator) []*ValidationError {
	var errors []*ValidationError

	for _, err := range err {
		var element ValidationError
		element.Key = strings.ToLower(err.Field())
		element.Message = err.Translate(translator)
		errors = append(errors, &element)
	}

	return errors
}
