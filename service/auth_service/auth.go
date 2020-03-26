package auth_service

import "github.com/namtx/go-gin-example/models"

type Auth struct {
	Username string
	Password string
}

func (auth *Auth) Check() (bool, error) {
	return models.CheckAuth(auth.Username, auth.Password)
}
