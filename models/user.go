package models

import "gorm.io/gorm"

// User contains the properties of a user.
type User struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Email    string
	Password string `json:"-"`
}
