package controllers

import (
	"log"
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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

	var Vehicles []models.Vehicle

	initializers.DB.Find(&Vehicles)

	c.JSON(200, gin.H{
		"vehicle": Vehicles,
	})
}

func VehicleInfoDelete(c *gin.Context) {
	var vehicle models.Vehicle
	id := c.Param(("id"))

	result := initializers.DB.Where("Id = ?", id).First(&vehicle)

	if result.Error != nil {
		log.Fatalf("cannot retrieve vehicle: %v\n", result.Error)
	}

	initializers.DB.Delete(&vehicle)

	c.JSON(200, gin.H{
		"result": "vehicle Deleted successsfully!",
	})
}

func VehicleInfoUpdate(c *gin.Context) {

	var vehicle models.Vehicle
	id := c.Param(("id"))

	var body struct {
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

func GetVehicleByUserId(c *gin.Context) {

	var vehicles []models.Vehicle
	Id := c.Param(("userid"))

	result := initializers.DB.Where("userid = ?", Id).Find(&vehicles)

	if result.Error != nil {
		log.Fatalf("cannot retrieve vehicle: %v\n", result.Error)
	}

	initializers.DB.Preload(clause.Associations).Where("userid = ?", Id).Find(&vehicles)

	c.JSON(200, gin.H{
		"vehicle": vehicles,
	})
}

func GetVehicleByVehId(c *gin.Context) {

	var vehicle models.Vehicle
	Id := c.Param(("id"))

	result := initializers.DB.Where("id = ?", Id).Find(&vehicle)

	if result.Error != nil {
		log.Fatalf("cannot retrieve vehicle: %v\n", result.Error)
	}

	initializers.DB.Preload(clause.Associations).Where("id = ?", Id).Find(&vehicle)

	c.JSON(200, gin.H{
		"vehicle": vehicle,
	})
}
