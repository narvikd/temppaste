package jsonreturn

import "github.com/gofiber/fiber/v2"

// Model contains the information necessary to create an API JSON response
type Model struct {
	Status  int
	Success bool
	Message string
	Data    interface{}
}

// NewModel returns a pointer to a new Model. Warning: "data" received as nil will be rewritten to an empty string.
func NewModel(status int, success bool, message string, data interface{}) *Model {
	if data == nil {
		data = ""
	}

	m := Model{
		Status:  status,
		Success: success,
		Message: message,
		Data:    data,
	}

	return &m
}

// Make makes a response from Model.
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
