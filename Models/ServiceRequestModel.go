package models

import (
	//"github.com/google/uuid"
	"time"

	"gorm.io/gorm"
)

type ServicesRequest struct {
	gorm.Model

	Id            string `gorm:"size:50;primaryKey"`
	Serviceid     uint
	Userid        uint
	Amount        float32
	ScheduledDate time.Time
	Users         User     `gorm:"foreignKey:Userid ;references:id;"`
	Services      Services `gorm:"foreignKey:Serviceid ;references:id;"`
}
