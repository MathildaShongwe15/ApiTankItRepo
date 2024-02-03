package main

import (
	controllers "myapp/Controllers"
	initializers "myapp/Initializers"
	middleware "myapp/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}
func main() {
	r := gin.Default()

	// router := gin.New()
	// router.Use(gin.Logger())

	//routes.AuthRoutes(router)
	//routes.UserRoutes(router)

	//r.POST("/createRequest", controllers.UserRequestCreate)
	//r.PUT("/request/:id", controllers.RequestTypeUpdate)
	//r.GET("/AllservicesRequested", controllers.ServiceTypeGetAll)

	//r.POST("/user", controllers.UserCreate)

	r.POST("/Auth", controllers.SignUp)
	r.POST("/Login", controllers.Login)
	r.GET("/Validate", middleware.RequireAuth, controllers.Validate)

	r.PUT("/ServiceUpdate", controllers.ServicesUpdate)
	r.GET("/AllServices", controllers.ServicesGetAll)
	r.POST("/ServiceCreate", controllers.ServiceCreate)
	r.DELETE("/DeleteService", controllers.ServicesDelete)

	r.PUT("/ServiceRequestUpdate")
	r.GET("/AllServiceRequests")
	r.POST("/ServiceRequestCreate", controllers.ServiceCreate)
	r.DELETE("/DeleteServiceRequest", controllers.ServicesDelete)
	r.Run()
}
