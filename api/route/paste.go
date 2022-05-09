package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"temppaste/api/route/model/createpaste"
	"temppaste/api/route/model/getpaste"
	"temppaste/internal/jsonreturn"
)

func (a *AppCtx) GetPaste(ginCtx *gin.Context) {
	p, err := getpaste.Get(ginCtx, a.DB)
	if err != nil {
		jsonreturn.Make(ginCtx, err)
		return
	}

	jsonreturn.OK(ginCtx, "paste retrieved successfully", p)
}

func (a *AppCtx) getPasteRaw(ginCtx *gin.Context) {
	p, err := getpaste.Get(ginCtx, a.DB)
	if err != nil {
		jsonreturn.Make(ginCtx, err)
		return
	}

	ginCtx.String(http.StatusOK, p.Content)
}

func (a *AppCtx) CreatePaste(ginCtx *gin.Context) {
	id, err := createpaste.Create(ginCtx, a.Translator, a.DB)
	if err != nil {
		jsonreturn.Make(ginCtx, err)
		return
	}

	jsonreturn.OK(ginCtx, "paste created", id)
}
