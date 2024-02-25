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
	r.GET("/Users", controllers.GetAllUsers)

	r.PUT("/ServiceUpdate/:id", controllers.ServicesUpdateById)
	r.GET("/AllServices", controllers.ServicesGetAll)
	r.POST("/ServiceCreate", controllers.ServiceCreate)
	r.DELETE("/DeleteService/:id", controllers.ServicesDeleteById)

	r.PUT("/ServiceRequestUpdate/:id", controllers.UserRequestUpdate)
	r.GET("/AllServiceRequests", controllers.UserRequestGetAll)
	r.POST("/ServiceRequestCreate", controllers.UserRequestCreate)
	r.DELETE("/DeleteServiceRequest/:id", controllers.UserRequestDelete)

	r.POST("/CreateVehicle", controllers.VehicleInfoCreate)
	r.GET("/GetAllVehicles", controllers.VehicleInfoGet)
	r.PUT("/UpdateVehicle/:id", controllers.VehicleInfoUpdate)
	r.DELETE("/DeleteVehicle/:id", controllers.VehicleInfoDelete)

	r.POST("/CreateProviders", controllers.ProvidersCreate)
	r.GET("/GetProviders", controllers.ProvidersGetAll)
	r.PUT("/UpdateProviderById/:id", controllers.ProvidersUpdateById)
	r.DELETE("/DeleteProviderById/:id", controllers.ProviderDeleteById)

	r.Run()
}
