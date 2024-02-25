package controllers

import (
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

}

func UserRequestUpdate(c *gin.Context) {

}
