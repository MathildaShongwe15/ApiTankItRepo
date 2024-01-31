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
	//initializers.DB.AutoMigrate(&models.ServicesRequest{})

}
