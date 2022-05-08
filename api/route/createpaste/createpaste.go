package createpaste

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-memdb"
	"temppaste/database/paste"
	"temppaste/internal/jsonreturn"
	"temppaste/internal/parser/customvalidator"
	"temppaste/internal/parser/fiberbodyparser"
	"temppaste/pkg/errorskit"
)

func Create(fiberCtx *fiber.Ctx, DB *memdb.MemDB) (string, *jsonreturn.Model) {
	model := new(paste.Paste)

	errBodyParser := fiberbodyparser.Parse(fiberCtx, model)
	if errBodyParser != nil {
		return "", jsonreturn.NewModel(
			fiber.StatusBadRequest, false, errBodyParser.Error(), "",
		)
	}
	errValidateStruct := customvalidator.Validate(model)
	for _, v := range errValidateStruct {
		return "", jsonreturn.NewModel(
			fiber.StatusBadRequest, false, v.Error(), "",
		)
	}

	id, err := paste.NewPaste(DB, model)
	if err != nil {
		errorskit.LogWrap(err, "createPaste endpoint")
		return "", jsonreturn.NewModel(
			fiber.StatusInternalServerError, false, "couldn't create paste", "",
		)
	}

	return id, nil
}
