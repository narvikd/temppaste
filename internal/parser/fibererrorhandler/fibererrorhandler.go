package fibererrorhandler

import (
	"github.com/gofiber/fiber/v2"
)

// Register registers a new Fiber Error Handler. It needs to be used with the Panic MW
//
// It overrides the REST error response when the application has an unexpected error
// (to avoid information disclosure to the client).
//
// Instead, it redirects the client to home.
//
// (More info on: https://docs.gofiber.io/api/middleware/recover and https://docs.gofiber.io/guide/error-handling)
//
func Register(ctx *fiber.Ctx) error {
	// Set Content-Type: text/plain; charset=utf-8
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message
	return ctx.Redirect("/")
}
