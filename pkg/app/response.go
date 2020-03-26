package app

import "github.com/namtx/go-gin-example/pkg/errors"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    httpCode,
		Message: errors.GetMessage(errCode),
		Data:    data,
	})

	return
}
