package route

import (
	"github.com/hashicorp/go-memdb"
	"temppaste/internal/app"
)

// AppCtx is a simple struct to include a collection of tools that a route could need to operate, for example a *memdb.MemDB.
type AppCtx struct {
	DB *memdb.MemDB
}

// newRouteCtx returns a pointer of a new instance of AppCtx.
func newRouteCtx(app *app.App) *AppCtx {
	routeCtx := AppCtx{
		DB: app.DB,
	}
	return &routeCtx
}

// Register registers fiber's routes.
func Register(app *app.App) {
	routes(app, newRouteCtx(app))
}

func routes(app *app.App, route *AppCtx) {
	app.Get("paste", route.getAll)
	app.Post("paste", route.newPaste)
}
