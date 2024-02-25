package controllers

import (
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
)

func CarInfoCreate(c *gin.Context) {

	var body struct {
		Id           string
		Userid       string
		VehicleBrand string
		VehicleModel string
		RegNo        string
		Color        string
		Description  string
	}

	c.Bind(&body)

	carInfo := models.Vehicle{Id: body.Id, Userid: body.Userid, VehicleBrand: body.VehicleBrand, VehicleModel: body.VehicleModel, RegNo: body.RegNo, Color: body.Color, Description: body.Description}
	result := initializers.DB.Create(&carInfo)

	//create a get
	if result.Error != nil {
		c.Status(400)
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
