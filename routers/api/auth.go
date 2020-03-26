package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/namtx/go-gin-example/pkg/app"
	"github.com/namtx/go-gin-example/pkg/errors"
	"github.com/namtx/go-gin-example/pkg/util"
	"github.com/namtx/go-gin-example/service/auth_service"
)

type authentication struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	ginApplication := app.Gin{C: c}
	validator := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")
	auth := authentication{Username: username, Password: password}

	ok, _ := validator.Valid(&auth)
	if !ok {
		app.LogErrors(validator.Errors)
		ginApplication.Response(http.StatusBadRequest, errors.InvalidParams, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		ginApplication.Response(http.StatusInternalServerError, errors.ErrorAuthCheckTokenFail, nil)
		return
	}
	if !isExist {
		ginApplication.Response(http.StatusUnauthorized, errors.ErrorAuth, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		ginApplication.Response(http.StatusInternalServerError, errors.ErrorIssueAuthTokenFail, nil)
		return
	}

	ginApplication.Response(http.StatusOK, errors.Success, map[string]string{"token": token})
}
