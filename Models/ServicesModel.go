package models

import (
	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type Services struct {
	gorm.Model
	Description string
	Type        string
}
