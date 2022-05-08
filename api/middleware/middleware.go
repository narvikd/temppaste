package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"temppaste/internal/app"
)

// InitMiddlewares initializes/registers all the app middlewares.
func InitMiddlewares(app *app.App) {
	initCorsMW(app)
	initRecoverMW(app)
}

// initCorsMW is set to allow all.
func initCorsMW(app *app.App) {
	app.Use(
		cors.New(cors.Config{
			AllowCredentials: true,
		}),
	)
}

// initRecoverMW initializes the Recover MW. If this is active, fibererrorhandler.Register also needs to be active
// to prevent information disclosure to the client.
func initRecoverMW(app *app.App) {
	app.Use(recover.New())
}
