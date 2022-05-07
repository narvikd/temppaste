package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/recover"
	"temppaste/internal/app"
)

// InitMiddlewares initializes/registers all the app middlewares.
func InitMiddlewares(app *app.App) {
	initRecoverMW(app)
}

// initRecoverMW initializes the Recover MW. If this is active, fibererrorhandler.Register also needs to be active
// to prevent information disclosure to the client.
func initRecoverMW(app *app.App) {
	app.Use(recover.New())
}
