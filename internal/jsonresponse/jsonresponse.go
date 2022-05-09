package jsonresponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Model contains the information necessary to create an API JSON response.
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
func Make(ctx *gin.Context, model *Model) {
	ctx.JSON(model.Status, gin.H{
		"success": model.Success,
		"message": model.Message,
		"data":    model.Data,
	})
}

// OK returns a successful response with status code 200
func OK(ctx *gin.Context, message string, data interface{}) {
	m := Model{
		Status:  http.StatusOK,
		Success: true,
		Message: message,
		Data:    data,
	}
	Make(ctx, &m)
}
