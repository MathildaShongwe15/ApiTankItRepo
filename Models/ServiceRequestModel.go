package models

import (
	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServicesRequest struct {
	gorm.Model

	Id        string `gorm:"size:50;primaryKey"`
	Serviceid uint
	Userid    uint
	Amount    float32
	Users     User     `gorm:"foreignKey:Userid ;references:id;"`
	Services  Services `gorm:"foreignKey:Serviceid ;references:id;"`
}
