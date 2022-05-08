package createpaste

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-memdb"
	"temppaste/database/paste"
	"temppaste/internal/fiberparser"
	"temppaste/internal/jsonreturn"
	"temppaste/pkg/errorskit"
)

func Create(fiberCtx *fiber.Ctx, DB *memdb.MemDB) (string, *jsonreturn.Model) {
	model := new(paste.Paste)

	errParseAndValidate := fiberparser.ParseAndValidate(fiberCtx, model)
	if errParseAndValidate != nil {
		return "", jsonreturn.NewModel(
			fiber.StatusBadRequest, false, errParseAndValidate.Error(), "",
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
