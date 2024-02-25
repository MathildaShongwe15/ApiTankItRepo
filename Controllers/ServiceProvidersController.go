package controllers

import (
	"log"
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
)

func ProvidersCreate(c *gin.Context) {

	var body struct {
		Id          string
		Name        string
		Email       string
		PhoneNumber string
	}

	c.Bind(&body)

	providers := models.ServiceProvider{Id: body.Id, Name: body.Name, Email: body.Email, PhoneNumber: body.PhoneNumber}
	result := initializers.DB.Create(&providers)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"providers": providers,
	})
}

func ProvidersGetAll(c *gin.Context) {

	var providers []models.ServiceProvider

	initializers.DB.Find(&providers)

	c.JSON(200, gin.H{
		"providers": providers,
	})
}

func ProviderDeleteById(c *gin.Context) {

	var provider models.ServiceProvider
	id := c.Param(("id"))

	result := initializers.DB.Where("Id = ?", id).First(&provider)
	if result.Error != nil {
		log.Fatalf("cannot retrieve Service Provider: %v\n", result.Error)
	}

	initializers.DB.Delete(&provider)

	c.JSON(200, gin.H{
		"result": "Service Provider Deleted successsfully!",
	})
}

func ProvidersUpdateById(c *gin.Context) {

	var provider models.ServiceProvider
	id := c.Param(("id"))

	var body struct {
		Name        string
		Email       string
		PhoneNumber string
	}

	c.Bind(&body)

	result := initializers.DB.Where("Id = ?", id).First(&provider)

	if result.Error != nil {
		log.Fatalf("cannot retrieve Service Provider: %v\n", result.Error)
	}

	initializers.DB.Model(&provider).Updates(models.ServiceProvider{
		Name:        body.Name,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
	})

	c.JSON(200, gin.H{
		"result": "Service Provider Updated successsfully!",
	})
}
