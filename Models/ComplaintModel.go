package models

import "gorm.io/gorm"

type Complaint struct {
	gorm.Model
	ServiceProviderId    string
	UserId               string
	ComplaintTitle       string `validate:"required, min=5, max=100"`
	ComplaintDescription string `validate:"required, min=5, max=100"`
	Status               string
	ServiceProvider      ServiceProvider `gorm:"foreignKey:ServiceProviderId ;references:Id;`
	User                 User            `gorm:"foreignKey:UserId ;references:Id;`
}
