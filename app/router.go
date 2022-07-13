package app

import (
	"github.com/julienschmidt/httprouter"
	"programmerzamannow/belajar-golang-restful-api/controller"
	"programmerzamannow/belajar-golang-restful-api/exception"
)

func NewRouter(categoryController controller.CategoryController, roleController controller.RoleController, userRoleController controller.UserRoleController, userProfileController controller.UserProfileController, userHotelController controller.UserHotelController, hotelController controller.HotelController, floorController controller.FloorController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/view-all-categories", categoryController.FindAll)
	router.GET("/api/view-categories-id/:categoryId", categoryController.FindById)
	router.POST("/api/create-categories", categoryController.Create)
	router.PUT("/api/update-categories-id/:categoryId", categoryController.Update)
	router.DELETE("/api/delete-categories-id/:categoryId", categoryController.Delete)

	router.GET("/api/view-all-role", roleController.FindAll)
	router.GET("/api/view-role-id/:roleId", roleController.FindById)
	router.POST("/api/create-role", roleController.Create)
	router.PUT("/api/update-role-id/:roleId", roleController.Update)
	router.DELETE("/api/delete-role-id/:roleId", roleController.Delete)

	router.GET("/api/view-all-user-role", userRoleController.FindAll)
	router.GET("/api/view-user-role-id/:userRoleId", userRoleController.FindById)
	router.POST("/api/create-user-role", userRoleController.Create)
	router.PUT("/api/update-user-role-id/:userRoleId", userRoleController.Update)
	router.DELETE("/api/delete-user-role-id/:userRoleId", userRoleController.Delete)

	router.GET("/api/view-all-user-profile", userProfileController.FindAll)
	router.GET("/api/view-user-profile-id/:userProfileId", userProfileController.FindById)
	router.POST("/api/create-user-profile", userProfileController.Create)
	router.PUT("/api/update-user-profile/:userProfileId", userProfileController.Update)
	router.DELETE("/api/delete-user-profile/:userProfileId", userProfileController.Delete)

	router.GET("/api/view-all-user-hotel", userHotelController.FindAll)
	router.GET("/api/view-user-hotel-id/:userHotelId", userHotelController.FindById)
	router.POST("/api/create-user-hotel", userHotelController.Create)
	router.PUT("/api/update-user-hotel-id/:userHotelId", userHotelController.Update)
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

	router.PanicHandler = exception.ErrorHandler
	return router
}
