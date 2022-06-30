package consts

import (
	"go_auth/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Create the JWT key used to create the signature
var JwtKey = []byte("my_secret_key")

var DB *gorm.DB

func InitializeDB() {
	var err error
	if DB != nil {
		return
	}
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	DB.AutoMigrate(&models.User{})
}
