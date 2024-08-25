package config

import (
	"github.com/chris097/gin-gorm-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Coonect() {
	db, err := gorm.Open(postgres.Open("host=127.0.0.1 user=postgres password=chris@1995 dbname=postgres port=5432"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	DB = db
}
