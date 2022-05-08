package customvalidator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// registerCustomValidators registers all custom validators for later use.
func registerCustomValidators(validate *validator.Validate, translator ut.Translator) {
	mapStructNamesToLowerCase(validate)
}

// mapStructNamesToLowerCase converts struct names like "Thing" to be displayed as "thing" on error output.
//
// If they contain an "-" it will be deleted "Thing-y" to "Thingy".
func mapStructNamesToLowerCase(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
