package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

type App struct {
	JwtSecret string
}

var AppSettings = &App{}

func Setup() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	AppSettings.JwtSecret = viper.GetString("JWT_SECRET")
}
