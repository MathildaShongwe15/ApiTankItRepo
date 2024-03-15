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

	r.POST("/Auth", controllers.SignUp)
	r.POST("/Login", controllers.Login)
	r.GET("/Validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/Users", controllers.GetAllUsers)
	r.GET("Users/:id", controllers.GetUserById)
	r.PUT("/ResetPassword/:email", controllers.ResetPassword)
	r.PUT("/UserUpdate/:id", controllers.UserUpdate)

	r.PUT("/ServiceUpdate/:id", controllers.ServicesUpdateById)
	r.GET("/AllServices", controllers.ServicesGetAll)
	r.GET(" ", controllers.GetServicesById)
	r.POST("/ServiceCreate", controllers.ServiceCreate)
	r.DELETE("/DeleteService/:id", controllers.ServicesDeleteById)

	r.PUT("/ServiceRequestUpdate/:id", controllers.UserRequestUpdate)
	r.GET("/AllServiceRequests", controllers.UserRequestGetAll)
	r.GET("/AllServiceRequestsById/:id", controllers.UserRequestGetById)
	r.POST("/ServiceRequestCreate", controllers.UserRequestCreate)
	r.DELETE("/DeleteServiceRequest/:id", controllers.UserRequestDelete)

	r.POST("/CreateVehicle", controllers.VehicleInfoCreate)
	r.GET("/GetAllVehicles", controllers.VehicleInfoGet)
	r.GET("/GetVehicleById", controllers.GetVehicleById)
	r.PUT("/UpdateVehicle/:id", controllers.VehicleInfoUpdate)
	r.DELETE("/DeleteVehicle/:id", controllers.VehicleInfoDelete)

	r.POST("/CreateProviders", controllers.ProvidersCreate)
	r.GET("/GetProviders", controllers.ProvidersGetAll)
	r.GET("/GetProviderById/:id", controllers.GetProviderById)
	r.PUT("/UpdateProviderById/:id", controllers.ProvidersUpdateById)
	r.DELETE("/DeleteProviderById/:id", controllers.ProviderDeleteById)
	r.GET("/GetProviderByService/:Serviceid", controllers.GetProviderByService)

	r.Run()
}
