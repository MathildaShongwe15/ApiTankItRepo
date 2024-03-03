package controllers

import (
	"log"
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
)

func VehicleInfoCreate(c *gin.Context) {

	var body struct {
		Id           string
		Userid       string
		VehicleBrand string
		VehicleModel string
		RegNo        string
		Color        string
		Description  string
	}

	c.BindJSON(&body)

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

func VehicleInfoGet(c *gin.Context) {

	var requests []models.ServicesRequest

	initializers.DB.Find(&requests)

	c.JSON(200, gin.H{
		"requests": requests,
	})
}

func VehicleInfoDelete(c *gin.Context) {
	var vehicle models.Vehicle
	id := c.Param(("id"))

	result := initializers.DB.Where("Id = ?", id).First(&vehicle)

	if result.Error != nil {
		log.Fatalf("cannot retrieve request: %v\n", result.Error)
	}

	initializers.DB.Delete(&vehicle)

	c.JSON(200, gin.H{
		"result": "Request Deleted successsfully!",
	})
}

func VehicleInfoUpdate(c *gin.Context) {

	var vehicle models.Vehicle
	id := c.Param(("id"))

	var body struct {
		Userid       string
		VehicleBrand string
		VehicleModel string
		RegNo        string
		Color        string
		Description  string
	}
	c.Bind(&body)

	result := initializers.DB.Where("Id = ?", id).First(&vehicle)

	if result.Error != nil {
		log.Fatalf("cannot retrieve request: %v\n", result.Error)
	}

	initializers.DB.Model(&vehicle).Updates(models.Vehicle{
		Userid:       body.Userid,
		VehicleBrand: body.VehicleBrand,
		VehicleModel: body.VehicleModel,
		RegNo:        body.RegNo,
		Color:        body.Color,
		Description:  body.Description,
	})

	c.JSON(200, gin.H{
		"result": " vehicle Updated successsfully!",
	})
}

func GetVehicleById(c *gin.Context) {

	var vehicle models.Vehicle
	id := c.Param(("id"))

	result := initializers.DB.Where("Id = ?", id).First(&vehicle)

	if result.Error != nil {
		log.Fatalf("cannot retrieve vehicle: %v\n", result.Error)
	}

	initializers.DB.Find(&vehicle)

	c.JSON(200, gin.H{
		"vehicle": vehicle,
	})
}