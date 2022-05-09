package createpaste

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/hashicorp/go-memdb"
	"github.com/narvikd/ginparser"
	"net/http"
	"temppaste/database/paste"
	"temppaste/internal/jsonreturn"
)

func Create(ginCtx *gin.Context, trans ut.Translator, DB *memdb.MemDB) (string, *jsonreturn.Model) {
	model := new(paste.Paste)
	errParse := ginparser.ParseAndValidate(ginCtx, trans, model)
	if errParse != nil {
		return "", jsonreturn.NewModel(http.StatusBadRequest, false, errParse.Error(), "")
	}

	id, err := paste.NewPaste(DB, model)
	if err != nil {
		return "", jsonreturn.NewModel(
			http.StatusInternalServerError, false, "couldn't create paste", "",
		)
	}
	return id, nil
}
