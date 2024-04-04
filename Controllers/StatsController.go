package controllers

import (
	"log"
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
)

func StatsCreate(c *gin.Context) {
	var body struct {
		ServiceProviderId string
		ReqPending        uint
		ReqCompleted      uint
		ReqCancelled      uint
		ReqLogged         uint
	}

	c.ShouldBindJSON(&body)

	values := models.Stats{ServiceProviderId: body.ServiceProviderId, ReqPending: body.ReqPending}
	result := initializers.DB.Create(&values)

	//create a get
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"values": values,
	})
}

func GetAllValuesByProviderId(c *gin.Context) {
	var stat models.Stats
	id := c.Param(("service_provider_id"))

	result := initializers.DB.Where("service_provider_id = ?", id).First(&stat)

	if result.Error != nil {
		log.Fatalf("cannot retrieve value: %v\n", result.Error)
	}

	initializers.DB.Where("service_provider_id = ?", id).First(&stat)

	c.JSON(200, gin.H{
		"values": stat,
	})
}

func UpdateStats(c *gin.Context) {
	var values models.Stats
	id := c.Param(("service_provider_id"))

	var body struct {
		ReqPending   uint
		ReqCompleted uint
		ReqCancelled uint
		ReqLogged    uint
	}

	c.ShouldBindJSON(&body)

	result := initializers.DB.Where("service_provider_id = ?", id).First(&values)

	if result.Error != nil {
		log.Fatalf("cannot retrieve values: %v\n", result.Error)
	}

	initializers.DB.Model(&values).Updates(models.Stats{
		ReqPending:   body.ReqPending,
		ReqCompleted: body.ReqCompleted,
		ReqCancelled: body.ReqCancelled,
		ReqLogged:    body.ReqLogged})

	c.JSON(200, gin.H{
		"values": values,
	})
}
