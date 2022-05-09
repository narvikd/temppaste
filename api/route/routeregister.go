package route

import (
	"embed"
	ut "github.com/go-playground/universal-translator"
	"github.com/hashicorp/go-memdb"
	"temppaste/internal/app"
	"temppaste/internal/ginembedfs"
)

// AppCtx is a simple struct to include a collection of tools that a route could need to operate, for example a *memdb.MemDB.
type AppCtx struct {
	DB           *memdb.MemDB
	PublicFolder embed.FS
	Translator   ut.Translator
}

// newRouteCtx returns a pointer of a new instance of AppCtx.
func newRouteCtx(app *app.App) *AppCtx {
	routeCtx := AppCtx{
		DB:           app.DB,
		PublicFolder: app.PublicFolder,
		Translator:   app.Translator,
	}
	return &routeCtx
}

// Register registers Gin's routes.
func Register(app *app.App) {
	routes(app, newRouteCtx(app))
}

func routes(app *app.App, route *AppCtx) {
	app.Use(
		ginembedfs.NewHandler("/", "public", app.PublicFolder),
	)
	app.GET("/p/:id", route.GetPaste)
	app.GET("/p/:id/raw", route.getPasteRaw)
	app.POST("/p", route.CreatePaste)
}
