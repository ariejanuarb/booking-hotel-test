package app

import (
	"github.com/julienschmidt/httprouter"
	"programmerzamannow/belajar-golang-restful-api/controller"
	"programmerzamannow/belajar-golang-restful-api/exception"
)

func NewRouter(categoryController controller.CategoryController, userProfileController controller.UserProfileController, userHotelController controller.UserHotelController, hotelController controller.HotelController, floorController controller.FloorController, employeeController controller.EmployeeController, roleController controller.RoleController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/view-all-categories", categoryController.FindAll)
	router.GET("/api/view-categories-id/:categoryId", categoryController.FindById)
	router.POST("/api/create-categories", categoryController.Create)
	router.PUT("/api/update-categories-id/:categoryId", categoryController.Update)
	router.DELETE("/api/delete-categories-id/:categoryId", categoryController.Delete)

	router.GET("/api/view-all-user-profile", userProfileController.FindAll)
	router.GET("/api/view-user-profile-id/:userProfileId", userProfileController.FindById)
	router.POST("/api/create-owner-profile", userProfileController.Create)
	router.PUT("/api/update-user-profile/:userProfileId", userProfileController.Update)
	router.DELETE("/api/delete-user-profile/:userProfileId", userProfileController.Delete)

	router.GET("/api/view-all-user-hotel", userHotelController.FindAll)
	router.GET("/api/view-user-hotel-id/:userHotelId", userHotelController.FindById)
	router.POST("/api/assign-employee-to-hotel", userHotelController.Create)
	router.PUT("/api/move-employee-to-another-hotel/:userHotelId", userHotelController.Update)
	router.DELETE("/api/delete-user-hotel-id/:userHotelId", userHotelController.Delete)

	router.GET("/api/view-all-hotel", hotelController.FindAll)
	router.GET("/api/view-hotel-id/:hotelId", hotelController.FindById)
	router.POST("/api/create-hotel", hotelController.Create)
	router.PUT("/api/update-hotel-id/:hotelId", hotelController.Update)
	router.DELETE("/api/delete-hotel-id/:hotelId", hotelController.Delete)

	router.GET("/api/view-all-floor", floorController.FindAll)
	router.GET("/api/view-floor-id/:floorId", floorController.FindById)
	router.POST("/api/create-floor", floorController.Create)
	router.PUT("/api/update-floor-id/:floorId", floorController.Update)
	router.DELETE("/api/delete-floor-id/:floorId", floorController.Delete)

	router.GET("/api/view-all-employee", employeeController.FindAll)
	router.GET("/api/view-employee-id/:employeeId", employeeController.FindById)
	router.POST("/api/create-new-employee", employeeController.Create)
	router.PUT("/api/update-employee-id/:employeeId", employeeController.Update)
	router.DELETE("/api/delete-employee-id/:employeeId", employeeController.Delete)

	router.GET("/api/view-all-role", roleController.FindAll)
	router.GET("/api/view-role-id/:roleId", roleController.FindById)
	router.POST("/api/create-new-role", roleController.Create)
	router.PUT("/api/update-role-id/:roleId", roleController.Update)
	router.DELETE("/api/delete-role-id/:roleId", roleController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
