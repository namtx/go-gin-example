package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/namtx/go-gin-example/routers/api"
)

func InitializeRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/auth", api.GetAuth)

	return r
}
