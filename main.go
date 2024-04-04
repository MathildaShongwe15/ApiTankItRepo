package main

import (
	controllers "myapp/Controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	//Mail "myapp/Mail"
	initializers "myapp/Initializers"
	middleware "myapp/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}
func main() {

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8081"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")
	r.Use(cors.New(corsConfig))

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
	r.PUT("/ServiceRequestUpdateStatus/:id", controllers.UserRequestUpdateStatus)
	r.GET("/AllServiceRequests", controllers.UserRequestGetAll)
	r.GET("/AllServiceRequestsById/:id", controllers.UserRequestGetById)
	r.GET("/UserRequestByProviderId/:service_provider_id", controllers.UserRequestByProviderId)
	r.POST("/ServiceRequestCreate", controllers.UserRequestCreate)
	r.DELETE("/DeleteServiceRequest/:id", controllers.UserRequestDelete)

	r.POST("/CreateVehicle", controllers.VehicleInfoCreate)
	r.GET("/GetAllVehicles", controllers.VehicleInfoGet)
	r.GET("/GetVehicleByUserId/:userid", controllers.GetVehicleByUserId)
	r.GET("/GetVehicleByVehicleId/:id", controllers.GetVehicleByVehId)
	r.PUT("/UpdateVehicle/:id", controllers.VehicleInfoUpdate)
	r.DELETE("/DeleteVehicle/:id", controllers.VehicleInfoDelete)

	r.POST("/CreateProviders", controllers.ProvidersCreate)
	r.GET("/GetProviders", controllers.ProvidersGetAll)
	r.GET("/GetProviderById/:id", controllers.GetProviderById)
	r.PUT("/UpdateProviderById/:id", controllers.ProvidersUpdateById)
	r.DELETE("/DeleteProviderById/:id", controllers.ProviderDeleteById)
	r.GET("/GetProviderByService/:Serviceid", controllers.GetProviderByService)

	r.POST("/CreateStats", controllers.StatsCreate)
	r.GET("/GetStatsById/:service_provider_id", controllers.GetAllValuesByProviderId)
	r.PUT("/UpdateStats/:service_provider_id", controllers.UpdateStats)

	r.Run()

}
