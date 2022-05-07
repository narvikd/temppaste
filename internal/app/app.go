package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-memdb"
)

// App is a simple struct to include a collection of tools that the webserver could need to operate.
//
// For example the pointer to DB that is used all over the application.
//
// This way the application can avoid the use of global variables.
type App struct {
	*fiber.App
	DB *memdb.MemDB
}
