package route

import (
	"github.com/gofiber/fiber/v2"
	"temppaste/database/paste"
	"temppaste/internal/jsonreturn"
)

func (a *AppCtx) getAll(fiberCtx *fiber.Ctx) error {
	pastes, err := paste.GetAllPastes(a.DB)
	if err != nil {
		return jsonreturn.BadRequest(fiberCtx, err.Error())
	}
	return jsonreturn.OK(fiberCtx, "success", pastes)
}
