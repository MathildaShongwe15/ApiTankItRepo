package models

import (
	"gorm.io/gorm"
)

type car struct {
	gorm.Model

	Id       string
	CarType  string
	RegNo    string
	Location string
}

//update location
