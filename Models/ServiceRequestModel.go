package models

import (
	"gorm.io/gorm"
)

type ServicesRequest struct {
	gorm.Model
	Id                string `gorm:"size:50;primaryKey"`
	Serviceid         uint
	Userid            string
	Vehicleid         string
	ServiceProviderId string
	Qauntity          string
	Type              string
	Spare             bool
	Amount            uint
	Longitude         float64
	Latitude          float64
	Accepted          bool
	ServiceProvider   ServiceProvider `gorm:"foreignKey:ServiceProviderId ;references:Id;`
	Users             User            `gorm:"foreignKey:Userid ;references:Id;"`
	Services          Services        `gorm:"foreignKey:Serviceid ;references:id;"`
	Vehicle           Vehicle         `gorm:"foreignKey:Vehicleid ;references:Id;"`
}
