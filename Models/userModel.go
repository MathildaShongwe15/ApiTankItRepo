package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"size:50"`
	Surname         string `gorm:"size:50"`
	Email           string `gorm:"size:100"`
	PhoneNumber     string `gorm:"size:15"`
	CarRegistration string `gorm:"size:15"`
	Cartype         string `gorm:"size:50"`
}
