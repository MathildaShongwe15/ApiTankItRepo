package controllers

import (
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
)

func ServiceCreate(c *gin.Context) {

	var body struct {
		Description string
		Type        string
	}

	c.Bind(&body)

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

// View all request types
func ServicesGetAll(c *gin.Context) {

	var services []models.Services

	initializers.DB.Find(&services)

	c.JSON(200, gin.H{
		"services": services,
	})
}

func ServicesDelete(c *gin.Context) {

	var services []models.Services
	id := c.Param(("id"))

	initializers.DB.Delete(&services, id)

	c.JSON(200, gin.H{
		"services": services,
	})
}

func ServicesUpdate(c *gin.Context) {

	id := c.Param(("id"))

	var body struct {
		Description string
		Type        string
	}

	c.Bind(&body)

	var service models.Services
	initializers.DB.First(&service, id)

	initializers.DB.Model(&service).Updates(models.Services{
		Description: body.Description,
		Type:        body.Type,
	})

	c.JSON(200, gin.H{
		"service": service,
	})
}
