package models

import "gorm.io/gorm"

type ServicesRequest struct {
	gorm.Model
	Id        string `gorm:"size:50;primaryKey"`
	Serviceid uint
	Userid    uint
	Vehicleid uint
	Qauntity  string
	Type      string
	Spare     bool
	Amount    float32
	Users     User     `gorm:"foreignKey:Userid ;references:Id;"`
	Services  Services `gorm:"foreignKey:Serviceid ;references:id;"`
	Vehicle   Vehicle  `gorm:"foreignKey:Vehicleid ;references:Id;"`
}
