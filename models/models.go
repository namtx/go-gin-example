package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func Setup() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "passw0rd")
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		logrus.Fatal("models.Setup err: ", err)
	}
}

func CloseDB() {
	defer db.Close()
}
