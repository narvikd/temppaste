package getpaste

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb"
	"net/http"
	"temppaste/database/paste"
	"temppaste/internal/jsonreturn"
)

func Get(ginCtx *gin.Context, DB *memdb.MemDB) (*paste.Paste, *jsonreturn.Model) {
	id := ginCtx.Param("id")
	if id == "" {
		return nil, jsonreturn.NewModel(
			http.StatusBadRequest, false, "id was empty", "",
		)
	}

	p, err := paste.GetPaste(DB, id)
	if err != nil {
		if err.Error() == "paste not found" {
			return nil, jsonreturn.NewModel(
				http.StatusNotFound, false, err.Error(), "",
			)
		}
		return nil, jsonreturn.NewModel(
			http.StatusInternalServerError, false, "couldn't get paste", "",
		)
	}

	return p, nil
}
