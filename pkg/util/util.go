package util

import "github.com/namtx/go-gin-example/pkg/setting"

func Setup() {
	jwtSecret = []byte(setting.AppSettings.JwtSecret)
}
