package models

import "gorm.io/gorm"

type Stats struct {
	gorm.Model
	ServiceProviderId string
	ReqPending        uint
	ReqCompleted      uint
	ReqCancelled      uint
	ReqLogged         uint
	ServiceProvider   ServiceProvider `gorm:"foreignKey:ServiceProviderId ;references:Id;`
}
