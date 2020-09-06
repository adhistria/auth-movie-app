package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	go_validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Validator represent validator entity
type Validator struct {
	validator  *go_validator.Validate
	translator ut.Translator
}

// NewValidator init validator
func NewValidator() *Validator {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")
	validate := go_validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	en_translations.RegisterDefaultTranslations(validate, trans)
	validator := Validator{
		validator:  validate,
		translator: trans,
	}
	return &validator
}

// Validate return error messages
func (v *Validator) Validate(data interface{}) map[string]string {
	err := v.validator.Struct(data)
	if err != nil {
		errors := map[string]string{}
		for _, e := range err.(go_validator.ValidationErrors) {
			errors[e.Field()] = e.Translate(v.translator)
		}
		return errors
	}
	return nil
}
