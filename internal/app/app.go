package app

import (
	"embed"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/hashicorp/go-memdb"
)

// App is a simple struct to include a collection of tools that the webserver could need to operate.
//
// For example the pointer to DB that is used all over the application.
//
// This way the application can avoid the use of global variables.
type App struct {
	*gin.Engine
	DB           *memdb.MemDB
	PublicFolder embed.FS
	Translator   ut.Translator
}
