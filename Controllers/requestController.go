package controllers

import (
	initializers "myapp/Initializers"
	models "myapp/Models"

	"github.com/gin-gonic/gin"
)

func RequestCreate(c *gin.Context) {

	//get data off req body

	var body struct {
		Type        string
		Description string
	}
	c.Bind(&body)

	request := models.Request{Type: body.Type, Description: body.Description}
	result := initializers.DB.Create(&request)

	//create a get
	if result.Error != nil {
		c.Status(400)
		return
	}

	//RETURN IT

	c.JSON(200, gin.H{
		"request": request,
	})
}

func RequestTypeGetAll(c *gin.Context) {

	var requests []models.Request
	//c.Bind((&body))

	initializers.DB.Find(&requests)

	// var requestTypes []models.RequestType
	// initializers.DB.Find(&requestTypes)

	//return the get
	c.JSON(200, gin.H{
		"requests": "requests",
	})

}
