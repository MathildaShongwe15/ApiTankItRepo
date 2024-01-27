package models

import (
	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRequest struct {
	gorm.Model

	Id        string `gorm:"size:50;primaryKey"`
	Requestid uint
	Userid    uint
	Amount    float32
	Requests  []Request `gorm:"foreignKey:id;references:Requestid"`
	Users     []User    `gorm:"foreignKey:id;references:Userid"`
}
