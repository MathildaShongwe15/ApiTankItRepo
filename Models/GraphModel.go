package models

type Graph struct {
	Id     string `gorm:"size:50;primaryKey"`
	XLabel string
	YCount uint
}
