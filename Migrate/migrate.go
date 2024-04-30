package main

import (
	initializers "myapp/Initializers"
	models "myapp/Models"
	//models "myapp/Models"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {

	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.ServicesRequest{})
	initializers.DB.AutoMigrate(&models.Services{})
	initializers.DB.AutoMigrate(&models.Vehicle{})
	initializers.DB.AutoMigrate(&models.ServiceProvider{})
	initializers.DB.AutoMigrate(&models.Stats{})
	initializers.DB.AutoMigrate(&models.Complaint{})

}
