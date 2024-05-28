package app

import (
	"github.com/gin-gonic/gin"

	"computesphere.com/computesphere-golang-rest-api-example/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data any) {
	g.C.JSON(
		httpCode,
		Response{
			Code: errCode,
			Msg:  e.GetMsg(errCode),
			Data: data,
		},
	)
	return
}
