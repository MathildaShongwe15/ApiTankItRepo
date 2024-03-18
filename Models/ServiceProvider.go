package models

type ServiceProvider struct {
	Id          string `gorm:"size:50;primaryKey"`
	Serviceid   uint
	Name        string
	Email       string
	PhoneNumber string `gorm:"size:10"`
	ServiceFee  uint
	Services    Services `gorm:"foreignKey:Serviceid ;references:id;"`
}
