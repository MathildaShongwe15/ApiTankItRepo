package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id          string `gorm:"size:50;primaryKey"`
	First_Name  string `validate:"required, min=5, max=100"`
	Last_Name   string `validate:"required, min=5, max=100"`
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"size:10"`
	Password    string
	Role        string ` validate:"required, eq=SERVICEPROVIDER|ew=CUSTOMER"`
}
