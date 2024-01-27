package models

import "gorm.io/gorm"

type Request struct {
	gorm.Model

	Type        string
	Description string
}
