package routes

import (
	"user-management/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine, path string) {

	// Initialize controller (dependency injection happens inside)
	direksiController := controller.NewDireksiController()

	// Group routes for Direksi
	direksiGroup := route.Group(path + "/direksi")
	{
		// Base CRUD operations
		direksiGroup.POST("/", direksiController.Create)
		direksiGroup.GET("/", direksiController.GetAll)
		direksiGroup.GET("/pagination", direksiController.GetPagination)
		direksiGroup.GET("/:uuid", direksiController.GetByUUID)
		direksiGroup.PUT("/:uuid", direksiController.Update)
		
		// Hard delete (permanent)
		direksiGroup.DELETE("/:uuid", direksiController.Delete)
		
		// Soft delete (mark as deleted)
		direksiGroup.DELETE("/:uuid/soft", direksiController.SoftDelete)
		
		// Search operations
		direksiGroup.GET("/search", direksiController.GetByField)
		direksiGroup.GET("/search/name", direksiController.FindByName)
		
		// Custom operations specific to Direksi
		direksiGroup.GET("/search/code", direksiController.GetByCode)
	}

	roleController := controller.NewRoleController()
	roleGroup := route.Group(path + "/role")
	{
		roleGroup.POST("/", roleController.Create)
		roleGroup.GET("/", roleController.GetAll)
		roleGroup.GET("/pagination", roleController.GetPagination)
		roleGroup.GET("/:uuid", roleController.GetByUUID)
		roleGroup.PUT("/:uuid", roleController.Update)
	}
}
