package models

import (
	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServicesRequest struct {
	gorm.Model

	Id        string `gorm:"size:50;primaryKey"`
	Requestid uint
	Userid    uint
	Amount    float32
	Type      string `gorm:"size:50;primaryKey"`
	Users     User   `gorm:"foreignKey:Userid ;references:id;"`
}
