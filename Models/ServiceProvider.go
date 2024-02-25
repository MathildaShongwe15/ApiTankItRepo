package models

import "gorm.io/gorm"

type ServiceProvider struct {
	gorm.Model
	Id          string `gorm:"size:50;primaryKey"`
	Name        string
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"size:10"`
}
