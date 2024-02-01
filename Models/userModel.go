package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email     string `gorm:"unique"`
	Password  string
	User_type *string `json:"user_type" validate:"required, eq=ADMIN|ew=USER"`
}
