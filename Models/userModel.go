package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	First_Name  *string `json:"first_name" validate:"required, min=5, max=100"`
	Last_Name   *string `json:"last_name" validate:"required, min=5, max=100"`
	Email       string  `gorm:"unique"`
	PhoneNumber string  `gorm:"size:15"`
	//Date_Of_Birth
	Country   string
	Password  string
	User_type *string `json:"user_type" validate:"required, eq=ADMIN|ew=USER"`
}
