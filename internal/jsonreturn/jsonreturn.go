package jsonreturn

import "github.com/gofiber/fiber/v2"

// Model contains the information necessary to create an API JSON response
type Model struct {
	Status  int
	Success bool
	Message string
	Data    interface{}
}

// Make makes a response from a Model's data
func Make(ctx *fiber.Ctx, model *Model) error {
	return ctx.Status(model.Status).JSON(&fiber.Map{
		"success": model.Success,
		"message": model.Message,
		"data":    model.Data,
	})
}

// OK returns a successful response with status code 200
func OK(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// NotFound returns a not found response with status code 404
func NotFound(ctx *fiber.Ctx, message string) error {
	return ctx.Status(404).JSON(&fiber.Map{
		"success": false,
		"message": message,
		"data":    "",
	})
}

// BadRequest returns a bad request response with status code http status 400
func BadRequest(ctx *fiber.Ctx, message string) error {
	return ctx.Status(400).JSON(&fiber.Map{
		"success": false,
		"message": message,
		"data":    "",
	})
}

// BadRequestWithData allows inserting data to return a bad request response with status code http status 400
func BadRequestWithData(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.Status(400).JSON(&fiber.Map{
		"success": false,
		"message": message,
		"data":    data,
	})
}

// DataNotValid returns a data not valid response with status code 422
func DataNotValid(ctx *fiber.Ctx, message string) error {
	return ctx.Status(422).JSON(&fiber.Map{
		"success": false,
		"message": message,
		"data":    "",
	})
}

// ServerError returns a server error response with status code 500
func ServerError(ctx *fiber.Ctx, message string) error {
	return ctx.Status(500).JSON(&fiber.Map{
		"success": false,
		"message": message,
		"data":    "",
	})
}

// TooManyRequests returns a response when the rate limiter is triggered
func TooManyRequests(ctx *fiber.Ctx, requestLimit string) error {
	msg := "Sorry, we're receiving too many requests from your IP address. Rate limited is " + requestLimit + "req/min"
	return ctx.Status(429).JSON(&fiber.Map{
		"success": false,
		"message": msg,
		"data":    "",
	})
}

// Unauthorised returns a Status Unauthorised status code (401)
func Unauthorised(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
		"success": false,
		"message": message,
		"data":    data,
	})
}
