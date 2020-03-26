package main

import (
	"net/http"

	"github.com/namtx/go-gin-example/models"
	"github.com/namtx/go-gin-example/pkg/setting"
	"github.com/namtx/go-gin-example/pkg/util"
	"github.com/namtx/go-gin-example/routers"
)

func init() {
	models.Setup()
	setting.Setup()
	util.Setup()
}

func main() {
	router := routers.InitializeRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
