package models

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model

	Userid   uint
	CarModel string `gorm:"size:100"`
	RegNo    string `gorm:"size:15"`
	Users    User   `gorm:"foreignKey:Userid ;references:id;"`
}
