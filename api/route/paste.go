package route

import (
	"github.com/gofiber/fiber/v2"
	"temppaste/api/route/model/createpaste"
	"temppaste/api/route/model/getpaste"
	"temppaste/internal/jsonreturn"
)

func (a *AppCtx) getPaste(fiberCtx *fiber.Ctx) error {
	p, err := getpaste.Get(fiberCtx, a.DB)
	if err != nil {
		return jsonreturn.Make(fiberCtx, err)
	}

	return jsonreturn.OK(fiberCtx, "paste retrieved successfully", p)
}

func (a *AppCtx) getPasteRaw(fiberCtx *fiber.Ctx) error {
	p, err := getpaste.Get(fiberCtx, a.DB)
	if err != nil {
		return jsonreturn.Make(fiberCtx, err)
	}

	return fiberCtx.Status(fiber.StatusOK).SendString(p.Content)
}

func (a *AppCtx) createPaste(fiberCtx *fiber.Ctx) error {
	id, err := createpaste.Create(fiberCtx, a.DB)
	if err != nil {
		return jsonreturn.Make(fiberCtx, err)
	}

	return jsonreturn.OK(fiberCtx, "paste created", id)
}
