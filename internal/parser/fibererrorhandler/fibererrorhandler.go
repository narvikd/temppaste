package fibererrorhandler

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// Register registers a new Fiber Error Handler. It needs to be used with the Panic MW
//
// It overrides the REST error response when the application has an unexpected error
// (to avoid information disclosure to the client).
//
// Instead, it just logs it, returning only "clientErrMsg".
//
// (More info on: https://docs.gofiber.io/api/middleware/recover and https://docs.gofiber.io/guide/error-handling)
//
func Register(ctx *fiber.Ctx, err error) error {
	const clientErrMsg = "500 Internal Server Error"
	code := fiber.StatusInternalServerError // Defaults to 500 status code

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message
	// TODO: Refactor
	log.Printf("Panic recovered: %v. %v\n", err, ctx)
	return ctx.Status(code).JSON(&fiber.Map{
		"success": false,
		"message": clientErrMsg,
		"data":    "",
	})
}
