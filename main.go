package main

import (
	"github.com/gin-gonic/gin"
	"github.com/narvikd/errorskit"
	"github.com/narvikd/ginparser"
	"github.com/narvikd/gorngseed"
	"log"
	"temppaste/api/route"
	"temppaste/internal/app"
	"temppaste/internal/database"
	"temppaste/internal/database/paste"
)

func main() {
	gorngseed.Register()
	db, errDBInit := database.NewDB(paste.NewSchema())
	if errDBInit != nil {
		log.Fatalln(errDBInit)
	}

	a := app.App{
		Engine:     gin.Default(),
		DB:         db,
		Translator: ginparser.Register(),
	}

	route.Register(&a)
	err := a.Run("0.0.0.0:3001")
	if err != nil {
		errorskit.FatalWrap(err, "couldn't start router")
	}
}
