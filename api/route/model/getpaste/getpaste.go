package getpaste

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb"
	"net/http"
	"temppaste/internal/database/paste"
	"temppaste/internal/jsonresponse"
)

func Get(ginCtx *gin.Context, DB *memdb.MemDB) (*paste.Paste, *jsonresponse.Model) {
	id := ginCtx.Param("id")
	if id == "" {
		return nil, jsonresponse.NewModel(
			http.StatusBadRequest, false, "id was empty", "",
		)
	}

	p, err := paste.GetPaste(DB, id)
	if err != nil {
		if err.Error() == "paste not found" {
			return nil, jsonresponse.NewModel(
				http.StatusNotFound, false, err.Error(), "",
			)
		}
		return nil, jsonresponse.NewModel(
			http.StatusInternalServerError, false, "couldn't get paste", "",
		)
	}

	return p, nil
}
