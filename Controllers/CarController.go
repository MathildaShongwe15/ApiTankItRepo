package controllers

import (
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
)

func CarInfoCreate(c *gin.Context) {

	var body struct {
		UserId   uint
		CarModel string
		RegNo    string
	}

	c.Bind(&body)

	carInfo := models.Car{CarModel: body.CarModel, RegNo: body.RegNo}
	result := initializers.DB.Create(&carInfo)

	//create a get
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"carInfo": carInfo,
	})

}

func CarInfoGet(c *gin.Context) {

	var requests []models.ServicesRequest

	initializers.DB.Find(&requests)

	c.JSON(200, gin.H{
		"requests": requests,
	})

}
