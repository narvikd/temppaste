package route

import (
	"embed"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/hashicorp/go-memdb"
	"net/http"
	"temppaste/internal/app"
)

// AppCtx is a simple struct to include a collection of tools that a route could need to operate, for example a *memdb.MemDB.
type AppCtx struct {
	DB           *memdb.MemDB
	PublicFolder embed.FS
}

// newRouteCtx returns a pointer of a new instance of AppCtx.
func newRouteCtx(app *app.App) *AppCtx {
	routeCtx := AppCtx{
		DB:           app.DB,
		PublicFolder: app.PublicFolder,
	}
	return &routeCtx
}

// Register registers fiber's routes.
func Register(app *app.App) {
	routes(app, newRouteCtx(app))
}

func routes(app *app.App, route *AppCtx) {
	loadEmbeddedHome(app)
	app.Get("paste/:id", route.getPaste)
	app.Post("paste", route.createPaste)
}

func loadEmbeddedHome(app *app.App) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(app.PublicFolder),
		Browse:       true,
		PathPrefix:   "/public",
		Index:        "index.html",
		NotFoundFile: "index.html",
		MaxAge:       3600,
	}))
}
