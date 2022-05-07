package route

import (
	"github.com/gofiber/fiber/v2"
	"temppaste/database/paste"
	"temppaste/internal/jsonreturn"
	"temppaste/internal/parser/customvalidator"
	"temppaste/internal/parser/fiberbodyparser"
	"temppaste/pkg/errorskit"
)

func (a *AppCtx) getPaste(fiberCtx *fiber.Ctx) error {
	id := fiberCtx.Params("id")
	if id == "" {
		return jsonreturn.BadRequest(fiberCtx, "id was empty")
	}

	p, err := paste.GetPaste(a.DB, id)
	if err != nil {
		if err.Error() == "paste not found" {
			return jsonreturn.NotFound(fiberCtx, err.Error())
		}

		errorskit.LogWrap(err, "get endpoint")
		return jsonreturn.ServerError(fiberCtx, "couldn't get paste")
	}
	return jsonreturn.OK(fiberCtx, "paste retrieved successfully", p)
}

func (a *AppCtx) createPaste(fiberCtx *fiber.Ctx) error {
	model := new(paste.Paste)

	errBodyParser := fiberbodyparser.Parse(fiberCtx, model)
	if errBodyParser != nil {
		return jsonreturn.BadRequest(fiberCtx, errBodyParser.Error())
	}
	errValidateStruct := customvalidator.Validate(model)
	for _, v := range errValidateStruct {
		return jsonreturn.BadRequest(fiberCtx, v.Error())
	}

	id, err := paste.NewPaste(a.DB, model)
	if err != nil {
		errorskit.LogWrap(err, "newPaste endpoint")
		return jsonreturn.BadRequest(fiberCtx, "couldn't create paste")
	}

	return jsonreturn.OK(fiberCtx, "paste created", id)
}
