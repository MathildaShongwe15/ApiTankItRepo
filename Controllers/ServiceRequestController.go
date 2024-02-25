package controllers

import (
	"log"
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
)

func UserRequestCreate(c *gin.Context) {

	var body struct {
		UserId     uint
		ServicesId uint
		Amount     float32
	}

	c.Bind(&body)

	serviceRequest := models.ServicesRequest{Userid: body.UserId, Serviceid: body.ServicesId, Amount: body.Amount}
	result := initializers.DB.Create(&serviceRequest)

	//create a get
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"request": serviceRequest,
	})
}

func UserRequestGetAll(c *gin.Context) {

	var requests []models.ServicesRequest

	initializers.DB.Find(&requests)

	c.JSON(200, gin.H{
		"requests": requests,
	})
}

func UserRequestDelete(c *gin.Context) {

	var request models.ServicesRequest
	id := c.Param(("id"))

	result := initializers.DB.Where("Id = ?", id).First(&request)
	if result.Error != nil {
		log.Fatalf("cannot retrieve request: %v\n", result.Error)
	}

	initializers.DB.Delete(&request)

	c.JSON(200, gin.H{
		"result": "Request Deleted successsfully!",
	})
}

func UserRequestUpdate(c *gin.Context) {
	var request models.ServicesRequest
	id := c.Param(("id"))

	var body struct {
		UserId     string
		ServicesId uint
		Amount     float32
	}

	c.Bind(&body)

	result := initializers.DB.Where("Id = ?", id).First(&request)

	if result.Error != nil {
		log.Fatalf("cannot retrieve request: %v\n", result.Error)
	}

	initializers.DB.Model(&request).Updates(models.ServicesRequest{
		Id:        body.UserId,
		Serviceid: body.ServicesId,
		Amount:    body.Amount,
	})

	c.JSON(200, gin.H{
		"result": " request Updated successsfully!",
	})
}
