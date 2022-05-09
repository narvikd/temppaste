package createpaste

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/hashicorp/go-memdb"
	"github.com/narvikd/ginparser"
	"net/http"
	"temppaste/internal/database/paste"
	"temppaste/internal/jsonresponse"
)

// Create handles JSON validations for the request, it returns a *jsonresponse.Model for easy return to the client.
func Create(ginCtx *gin.Context, trans ut.Translator, DB *memdb.MemDB) (string, *jsonresponse.Model) {
	model := new(paste.Paste)
	errParse := ginparser.ParseAndValidate(ginCtx, trans, model)
	if errParse != nil {
		return "", jsonresponse.NewModel(http.StatusBadRequest, false, errParse.Error(), "")
	}

	id, err := paste.NewPaste(DB, model)
	if err != nil {
		return "", jsonresponse.NewModel(
			http.StatusInternalServerError, false, "couldn't create paste", "",
		)
	}
	return id, nil
}
