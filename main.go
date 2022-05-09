package main

import (
	"github.com/gin-gonic/gin"
	"github.com/narvikd/ginparser"
	"github.com/narvikd/gorngseed"
	"log"
	"temppaste/api/route"
	"temppaste/database"
	"temppaste/database/paste"
	"temppaste/internal/app"
	"temppaste/pkg/errorskit"
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
		errorskit.LogWrap(err, "couldn't start router")
	}
}
