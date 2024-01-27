package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model

	Id        string
	Location  string
	TotalCost float32
}
