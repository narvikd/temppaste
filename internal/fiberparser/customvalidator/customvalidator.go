package customvalidator

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslation "github.com/go-playground/validator/v10/translations/en"
)

// ValidateJSON accepts a struct and validates it according to its model's tags. It returns a slice of errors.
//
// When iterating there's no need to check if errors are nil, because it doesn't return empty errors.
func ValidateJSON(s interface{}) []error {
	validate, translator := createNewValidation()
	registerCustomValidators(validate, translator)

	// errStructValidation holds all the possible errors that could be caused when doing a new validation on a struct.
	errStructValidation := validate.Struct(s)
	return translateErrors(errStructValidation, translator)
}

// createNewValidation creates the necessary components in order to invoque a new validation
func createNewValidation() (*validator.Validate, ut.Translator) {
	validate := validator.New()
	english := en.New()
	uniTranslator := ut.New(english, english)
	translator, _ := uniTranslator.GetTranslator("en")
	_ = entranslation.RegisterDefaultTranslations(validate, translator) // Can safely ignore this error
	return validate, translator
}

// translateErrors takes wrapped errors, translates them, and returns them as a normal error slice.
func translateErrors(err error, trans ut.Translator) (errs []error) {
	// If there aren't any errors, return an empty slice.
	if err == nil {
		return nil
	}

	// For each error, translate it and append it.
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
