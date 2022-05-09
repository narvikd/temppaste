package ginparser

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslation "github.com/go-playground/validator/v10/translations/en"
)

func Register() ut.Translator {
	validate, translator := createNewValidation()
	registerCustomValidators(validate, translator)
	return translator
}

// createNewValidation creates the necessary components in order to invoque a new validation
func createNewValidation() (*validator.Validate, ut.Translator) {
	validate := binding.Validator.Engine().(*validator.Validate)
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
	validatorErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		if err.Error() == "EOF" {
			errs = append(errs, errors.New("no valid data was sent to the server"))
		}
		errs = append(errs, err)
		return errs
	}
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
