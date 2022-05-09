package ginparser

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

func ParseAndValidate(ctx *gin.Context, trans ut.Translator, s interface{}) error {
	errBind := ctx.ShouldBindJSON(s)
	if errBind != nil {
		errs := translateErrors(errBind, trans)
		for _, v := range errs {
			return v
		}
	}
	return nil
}
