package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-memdb"
	"github.com/narvikd/fiberparser"
	"log"
	"temppaste/api/middleware"
	"temppaste/api/route"
	"temppaste/database"
	"temppaste/database/paste"
	"temppaste/internal/app"
	"temppaste/internal/app/shutdown"
	"temppaste/pkg/errorskit"
	"time"
)

// startFiber starts Fiber on a separate thread that blocks.
//
// Using a different thread is needed to make subsequent functions to work.
//
// It also registers the application shutdown, so it can exit cleanly when "Ctrl + C" is pressed.
func startFiber(publicFolder embed.FS) {
	db, errDBInit := database.NewDB(paste.NewSchema())
	if errDBInit != nil {
		log.Fatalln(errDBInit)
	}

	a := newApp(db, publicFolder)

	go func() {
		errFiberStart := a.Listen("0.0.0.0:3001")
		if errFiberStart != nil {
			log.Fatalln(errorskit.Wrap(errFiberStart, "server can't be started"))
		}
	}()

	shutdown.Register(a)
}

// newApp returns an instead of a new app.App pointer.
func newApp(db *memdb.MemDB, publicFolder embed.FS) *app.App {
	a := app.App{
		App: fiber.New(fiber.Config{
			AppName:           "Temp Paste",
			EnablePrintRoutes: false,
			IdleTimeout:       time.Second * 5, // Max time to wait for the next request when keep-alive is enabled.
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				return fiberparser.RegisterErrorHandler(ctx)
			},
		}),
		DB:           db,
		PublicFolder: publicFolder,
	}

	middleware.InitMiddlewares(&a)
	route.Register(&a)
	return &a
}
