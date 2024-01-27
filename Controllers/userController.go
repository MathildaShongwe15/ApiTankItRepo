package controllers

import (
	// initializers "myapp/Initializers"
	// models "myapp/Models"

	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {

	//get data off req body

	var body struct {
		ID              uint   `gorm:"primaryKey"`
		Name            string `gorm:"size:50"`
		Surname         string `gorm:"size:50"`
		Email           string `gorm:"size:100"`
		PhoneNumber     string `gorm:"size:15"`
		CarRegistration string `gorm:"size:15"`
		Cartype         string `gorm:"size:50"`
	}
	c.Bind(&body)

	user := models.Users{Name: body.Name, Surname: body.Surname, Email: body.Email, PhoneNumber: body.PhoneNumber, CarRegistration: body.CarRegistration, Cartype: body.Cartype}
	result := initializers.DB.Create(&user)

	//create a get
	if result.Error != nil {
		c.Status(400)
		return
	}

	//RETURN IT

	c.JSON(200, gin.H{
		"user": user,
	})
}
