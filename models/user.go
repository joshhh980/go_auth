package models

import (
	"go_auth/responses"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Email    string
	Password string
	Name     string
}

func (u User) BuildUser() responses.UserResponse {
	return responses.UserResponse{
		ID:    u.ID,
		Email: u.Email,
	}
}
