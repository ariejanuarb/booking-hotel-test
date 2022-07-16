package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"programmerzamannow/belajar-golang-restful-api/app"
	"programmerzamannow/belajar-golang-restful-api/controller"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/middleware"
	"programmerzamannow/belajar-golang-restful-api/repository"
	"programmerzamannow/belajar-golang-restful-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	userProfileRepository := repository.NewUserProfileRepository()
	userProfileService := service.NewUserProfileService(userProfileRepository, db, validate)
	userProfileController := controller.NewUserProfileController(userProfileService)

	userHotelRepository := repository.NewUserHotelRepository()
	userHotelService := service.NewUserHotelService(userHotelRepository, db, validate)
	userHotelController := controller.NewUserHotelController(userHotelService)

	hotelRepository := repository.NewHotelRepository()
	hotelService := service.NewHotelService(hotelRepository, db, validate)
	hotelController := controller.NewHotelController(hotelService)

	floorRepository := repository.NewFloorRepository()
	floorService := service.NewFloorService(floorRepository, db, validate)
	floorController := controller.NewFloorController(floorService)

	employeeRepository := repository.NewEmployeeRepository()
	employeeService := service.NewEmployeeService(employeeRepository, db, validate)
	employeeController := controller.NewEmployeeController(employeeService)

	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	router := app.NewRouter(categoryController, userProfileController, userHotelController, hotelController, floorController, employeeController, roleController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
