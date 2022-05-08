package getpaste

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-memdb"
	"temppaste/database/paste"
	"temppaste/internal/jsonreturn"
	"temppaste/pkg/errorskit"
)

func Get(fiberCtx *fiber.Ctx, DB *memdb.MemDB) (*paste.Paste, *jsonreturn.Model) {
	id := fiberCtx.Params("id")
	if id == "" {
		return nil, jsonreturn.NewModel(
			fiber.StatusBadRequest, false, "id was empty", "",
		)
	}

	p, err := paste.GetPaste(DB, id)
	if err != nil {
		if err.Error() == "paste not found" {
			return nil, jsonreturn.NewModel(
				fiber.StatusNotFound, false, err.Error(), "",
			)
		}
		errorskit.LogWrap(err, "getPaste endpoint")

		return nil, jsonreturn.NewModel(
			fiber.StatusInternalServerError, false, "couldn't get paste", "",
		)
	}

	return p, nil
}
