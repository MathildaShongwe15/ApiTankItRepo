package models

import (
	"gorm.io/gorm"
)

type Ratings struct {
	gorm.Model

	Id        string
	Userid    uint
	Serviceid uint
	upvotes   uint
	feedback  string
	Users     User     `gorm:"foreignKey:Userid ;references:id;"`
	Services  Services `gorm:"foreignKey:Serviceid ;references:id;"`
}
