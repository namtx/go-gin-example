package models

import (
	"github.com/jinzhu/gorm"
)

type Auth struct {
	gorm.Model

	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, error) {
	var auth Auth

	a := &Auth{
		Username: username,
		Password: password,
	}

	err := db.Select("id").Where(a).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
