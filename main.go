package main

import (
	controllers "myapp/Controllers"
	initializers "myapp/Initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}
func main() {
	r := gin.Default()
	r.POST("/request", controllers.RequestCreate)
	//r.PUT("/request/:id", controllers.RequestTypeUpdate)
	r.GET("/requestAvailable", controllers.RequestTypeGetAll)
	r.GET("/request:id", controllers.FindRequest)

	r.POST("/user", controllers.UserCreate)
	r.Run()
}
