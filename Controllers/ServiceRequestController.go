package controllers

import (
	"log"
	initializers "myapp/Initializers"
	models "myapp/Models"
	"net/smtp"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func UserRequestCreate(c *gin.Context) {

	var body struct {
		Id                string
		Serviceid         uint
		Userid            string
		Vehicleid         string
		ServiceProviderId string
		Qauntity          string
		Type              string
		Spare             bool
		Amount            uint
		Accepted          bool
		Longitude         float64
		Latitude          float64
	}

	c.ShouldBindJSON(&body)

	serviceRequest := models.ServicesRequest{Id: body.Id, Serviceid: body.Serviceid, Userid: body.Userid, Vehicleid: body.Vehicleid, ServiceProviderId: body.ServiceProviderId, Qauntity: body.Qauntity, Type: body.Type, Spare: body.Spare, Amount: body.Amount, Accepted: body.Accepted, Longitude: body.Longitude, Latitude: body.Latitude}
	result := initializers.DB.Create(&serviceRequest)

	if result.Error != nil {
		c.Status(400)
		return
	}

	RequestEmail()
	c.JSON(200, gin.H{
		"request": serviceRequest,
	})
}
func UserRequestGetAll(c *gin.Context) {

	var requests []models.ServicesRequest

	initializers.DB.Preload(clause.Associations).Find(&requests)

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
		Longitude float64
		Latitude  float64
	}

	c.ShouldBindJSON(&body)

	result := initializers.DB.Where("Id = ?", id).First(&request)

	if result.Error != nil {
		log.Fatalf("cannot retrieve request: %v\n", result.Error)
	}

	initializers.DB.Model(&request).Updates(models.ServicesRequest{

		Longitude: body.Longitude,
		Latitude:  body.Latitude,
	})

	c.JSON(200, gin.H{
		"result": request,
	})
}
func UserRequestGetById(c *gin.Context) {

	var request models.ServicesRequest
	id := c.Param(("id"))

	result := initializers.DB.Where("id = ?", id).First(&request)

	if result.Error != nil {
		log.Fatalf("cannot retrieve request: %v\n", result.Error)
	}

	initializers.DB.Preload(clause.Associations).Find(&request)

	c.JSON(200, gin.H{
		"requests": request,
	})
}
func UserRequestGetByUserId(c *gin.Context) {

	var request models.ServicesRequest
	id := c.Param(("userid"))

	result := initializers.DB.Where("userid = ?", id).First(&request)

	if result.Error != nil {
		log.Fatalf("cannot retrieve request: %v\n", result.Error)
	}

	initializers.DB.Preload(clause.Associations).Find(&request)

	c.JSON(200, gin.H{
		"requests": request,
	})
}
func UserRequestByProviderId(c *gin.Context) {

	var requests []models.ServicesRequest
	providerId := c.Param(("service_provider_id"))

	result := initializers.DB.Where("service_provider_id = ?", providerId).Find(&requests)

	if result.Error != nil {
		log.Fatalf("cannot retrieve request: %v\n", result.Error)
	}

	initializers.DB.Preload(clause.Associations).Where("service_provider_id = ?", providerId).Find(&requests)
	c.JSON(200, gin.H{
		"requests": requests,
	})
}
func UserRequestUpdateStatus(c *gin.Context) {
	var request models.ServicesRequest
	id := c.Param(("id"))

	var body struct {
		Accepted bool
	}

	c.ShouldBindJSON(&body)

	result := initializers.DB.Where("Id = ?", id).First(&request)

	if result.Error != nil {
		log.Fatalf("cannot retrieve request: %v\n", result.Error)
	}

	initializers.DB.Model(&request).Updates(models.ServicesRequest{

		Accepted: body.Accepted,
	})

	c.JSON(200, gin.H{
		"result": request,
	})
}
func RequestEmail() {
	auth := smtp.PlainAuth(
		"",
		"tankitroadsideassistance@gmail.com",
		"mflqvpvhtjfvbevg",
		"smtp.gmail.com",
	)
	msg := "Subject: Request has been Placed\nYou have successfully placed your Request. If this was not you please report to tankitroadsideassistance@gmail.com "

	smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"tankitroadsideassistance@gmail.com",
		[]string{"tankitroadsideassistance@gmail.com"},
		[]byte(msg),
	)

}
