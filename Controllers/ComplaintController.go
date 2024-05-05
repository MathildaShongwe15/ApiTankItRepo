package controllers

import (
	"log"
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CreateComplaint(c *gin.Context) {

	var body struct {
		ServiceProviderId    string
		UserId               string
		ComplaintTitle       string
		ComplaintDescription string
	}

	c.BindJSON(&body)

	complaint := models.Complaint{ServiceProviderId: body.ServiceProviderId, UserId: body.UserId, ComplaintTitle: body.ComplaintTitle, ComplaintDescription: body.ComplaintDescription}
	result := initializers.DB.Create(&complaint)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"complaint": complaint,
	})
}

func GetAllComplaints(c *gin.Context) {

	var complaints []models.Complaint

	initializers.DB.Find(&complaints)

	initializers.DB.Preload(clause.Associations).Find(&complaints)

	c.JSON(200, gin.H{
		"Complaints": complaints,
	})
}

func GetComplaintsByProviderId(c *gin.Context) {

	var complaint models.Complaint
	id := c.Param(("service_provider_id"))

	result := initializers.DB.Where("service_provider_id = ?", id).First(&complaint)

	if result.Error != nil {
		log.Fatalf("cannot retrieve Complaints: %v\n", result.Error)
	}

	initializers.DB.Find(&complaint)

	c.JSON(200, gin.H{
		"Complaints": complaint,
	})
}

func UpdateComplaintsById(c *gin.Context) {

	var Complaints models.Complaint
	id := c.Param(("id"))

	var body struct {
		ServiceProviderId    string
		UserId               string
		ComplaintTitle       string
		ComplaintDescription string
	}

	c.Bind(&body)

	result := initializers.DB.Where("Id = ?", id).First(&Complaints)

	if result.Error != nil {
		log.Fatalf("cannot retrieve Service Provider: %v\n", result.Error)
	}

	initializers.DB.Model(&Complaints).Updates(models.Complaint{
		ServiceProviderId:    body.ServiceProviderId,
		UserId:               body.UserId,
		ComplaintTitle:       body.ComplaintTitle,
		ComplaintDescription: body.ComplaintDescription,
	})

	c.JSON(200, gin.H{
		"result": "Complaint Updated successsfully!",
	})
}

func DeleteComplaintsById(c *gin.Context) {

	var complaints models.Complaint
	id := c.Param(("id"))

	result := initializers.DB.Where("Id = ?", id).First(&complaints)
	if result.Error != nil {
		log.Fatalf("cannot retrieve complaints: %v\n", result.Error)
	}

	initializers.DB.Delete(&complaints)

	c.JSON(200, gin.H{
		"result": "Complaints Deleted successsfully!",
	})
}
