package controllers

import (
	"log"
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func ProvidersCreate(c *gin.Context) {

	var body struct {
		Id          string
		ServiceId   uint
		Name        string
		Email       string
		PhoneNumber string
		ServiceFee  uint
	}

	c.BindJSON(&body)

	providers := models.ServiceProvider{Id: body.Id, Serviceid: body.ServiceId, Name: body.Name, Email: body.Email, PhoneNumber: body.PhoneNumber, ServiceFee: body.ServiceFee}
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

	initializers.DB.Preload(clause.Associations).Find(&providers)

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

func GetProviderById(c *gin.Context) {

	var provider models.ServiceProvider
	id := c.Param(("id"))

	result := initializers.DB.Where("Id = ?", id).First(&provider)

	if result.Error != nil {
		log.Fatalf("cannot retrieve provider: %v\n", result.Error)
	}

	initializers.DB.Find(&provider)

	c.JSON(200, gin.H{
		"provider": provider,
	})
}

func GetProviderByService(c *gin.Context) {

	var providers []models.ServiceProvider

	serviceid := c.Param(("Serviceid"))

	result := initializers.DB.Where("Serviceid = ?", serviceid).Find(&providers)

	if result.Error != nil {
		log.Fatalf("cannot retrieve providers: %v\n", result.Error)
	}

	initializers.DB.Where("Serviceid = ?", serviceid).Find(&providers)

	c.JSON(200, gin.H{
		"providers": providers,
	})
}
