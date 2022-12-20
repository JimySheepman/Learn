package main

import (
	"github.com/JimySheepman/class/controllers"
	"github.com/JimySheepman/class/dto"
	"github.com/JimySheepman/class/repository"
	"github.com/JimySheepman/class/services"
)

func main() {

	// repository
	adminRepository := repository.NewAdminRepository()

	// services
	adminServices := services.NewAdminServices(adminRepository)

	adminController := controllers.NewAdminControllers(adminServices)
	asyncAdminController := controllers.NewAsyncAdminControllers(adminServices)

	adminController.ReadHandler(&dto.RestRequest{})
	asyncAdminController.ReadHandler2(&dto.RestRequest2{})
}
