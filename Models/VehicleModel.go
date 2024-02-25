package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model

	Id           string `gorm:"size:50;primaryKey"`
	Userid       string `gorm:"size:50"`
	VehicleBrand string `gorm:"size:50"`
	VehicleModel string `gorm:"size:50"`
	RegNo        string `gorm:"size:10"`
	Color        string `gorm:"size:50"`
	Description  string `gorm:"size:100"`
	Users        User   `gorm:"foreignKey:Userid ;references:Id;"`
}
