package route

import (
	"github.com/gofiber/fiber/v2"
	"temppaste/database/paste"
	"temppaste/internal/jsonreturn"
	"temppaste/internal/parser/customvalidator"
	"temppaste/internal/parser/fiberbodyparser"
)

func (a *AppCtx) newPaste(fiberCtx *fiber.Ctx) error {
	model := new(paste.Paste)

	errBodyParser := fiberbodyparser.Parse(fiberCtx, model)
	if errBodyParser != nil {
		return jsonreturn.BadRequest(fiberCtx, errBodyParser.Error())
	}
	errValidateStruct := customvalidator.Validate(model)
	for _, v := range errValidateStruct {
		return jsonreturn.BadRequest(fiberCtx, v.Error())
	}

	err := paste.NewPaste(a.DB, model)
	if err != nil {
		return jsonreturn.BadRequest(fiberCtx, err.Error())
	}

	return jsonreturn.OK(fiberCtx, "success", nil)
}
