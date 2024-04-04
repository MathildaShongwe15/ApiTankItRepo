package controllers

import (
	"log"
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
)

func ServiceCreate(c *gin.Context) {

	var body struct {
		Description string
		Type        string
	}

	c.BindJSON(&body)

	services := models.Services{Description: body.Description, Type: body.Type}
	result := initializers.DB.Create(&services)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"services": services,
	})
}

func ServicesGetAll(c *gin.Context) {

	var services []models.Services

	initializers.DB.Find(&services)

	c.JSON(200, gin.H{
		"services": services,
	})
}

func ServicesDeleteById(c *gin.Context) {

	var services models.Services
	id := c.Param(("id"))

	result := initializers.DB.Where("Id = ?", id).First(&services)
	if result.Error != nil {
		log.Fatalf("cannot retrieve service: %v\n", result.Error)
	}

	initializers.DB.Delete(&services)

	c.JSON(200, gin.H{
		"result": "Service Deleted successsfully!",
	})
}

func ServicesUpdateById(c *gin.Context) {

	var service models.Services
	id := c.Param(("id"))

	var body struct {
		Description string
		Type        string
	}

	c.Bind(&body)

	result := initializers.DB.Where("Id = ?", id).First(&service)

	if result.Error != nil {
		log.Fatalf("cannot retrieve Service: %v\n", result.Error)
	}

	initializers.DB.Model(&service).Updates(models.Services{
		Description: body.Description,
		Type:        body.Type,
	})

	c.JSON(200, gin.H{
		"service": "Service  Updated successsfully!",
	})
}

func GetServicesById(c *gin.Context) {

	var service models.Services
	id := c.Param(("id"))

	result := initializers.DB.Where("Id = ?", id).First(&service)

	if result.Error != nil {
		log.Fatalf("cannot retrieve service: %v\n", result.Error)
	}

	initializers.DB.Find(&service)

	c.JSON(200, gin.H{
		"service": service,
	})
}
