package baseController

import (
	"strconv"
	baseService "user-management/service/base"
	"user-management/util"

	"github.com/gin-gonic/gin"
)

type BaseControllerInterface[T any] interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByUUID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	SoftDelete(c *gin.Context)
	GetPagination(c *gin.Context)
	GetByField(c *gin.Context)
	FindByName(c *gin.Context)
}

type BaseController[T any] struct {
	Service baseService.BaseServiceInterface[T]
}

func NewBaseController[T any](service baseService.BaseServiceInterface[T]) *BaseController[T] {
	return &BaseController[T]{Service: service}
}

func (ctrl *BaseController[T]) Create(c *gin.Context) {
	var entity T
	if err := c.ShouldBindJSON(&entity); err != nil {
		util.ErrorBadRequest(c, "Invalid request body: "+err.Error())
		return
	}

	if err := ctrl.Service.Create(&entity); err != nil {
		util.ErrorInternalServer(c, "Failed to create entity: "+err.Error())
		return
	}

	util.SuccessCreated(c, "Entity created successfully", entity)
}

func (ctrl *BaseController[T]) GetAll(c *gin.Context) {
	var entities []T
	if err := ctrl.Service.GetAll(&entities); err != nil {
		util.ErrorInternalServer(c, "Failed to get entities: "+err.Error())
		return
	}

	util.SuccessOK(c, "Entities retrieved successfully", entities)
}

func (ctrl *BaseController[T]) GetByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		util.ErrorBadRequest(c, "UUID parameter is required")
		return
	}

	var entity T
	if err := ctrl.Service.GetByUUID(&entity, uuid); err != nil {
		util.ErrorNotFound(c, "Entity not found: "+err.Error())
		return
	}

	util.SuccessOK(c, "Entity retrieved successfully", entity)
}

func (ctrl *BaseController[T]) Update(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		util.ErrorBadRequest(c, "UUID parameter is required")
		return
	}

	var entity T
	if err := c.ShouldBindJSON(&entity); err != nil {
		util.ErrorBadRequest(c, "Invalid request body: "+err.Error())
		return
	}

	if err := ctrl.Service.Update(&entity, uuid); err != nil {
		util.ErrorInternalServer(c, "Failed to update entity: "+err.Error())
		return
	}

	util.SuccessOK(c, "Entity updated successfully", entity)
}

func (ctrl *BaseController[T]) Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		util.ErrorBadRequest(c, "UUID parameter is required")
		return
	}

	if err := ctrl.Service.Delete(uuid); err != nil {
		util.ErrorInternalServer(c, "Failed to delete entity: "+err.Error())
		return
	}

	util.SuccessOK(c, "Entity deleted successfully", nil)
}

func (ctrl *BaseController[T]) SoftDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		util.ErrorBadRequest(c, "UUID parameter is required")
		return
	}

	if err := ctrl.Service.SoftDelete(uuid); err != nil {
		util.ErrorInternalServer(c, "Failed to soft delete entity: "+err.Error())
		return
	}

	util.SuccessOK(c, "Entity soft deleted successfully", nil)
}

func (ctrl *BaseController[T]) GetPagination(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		util.ErrorBadRequest(c, "Invalid page parameter")
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		util.ErrorBadRequest(c, "Invalid page_size parameter")
		return
	}

	var entities []T
	pagination, err := ctrl.Service.GetPagination(page, pageSize, &entities)
	if err != nil {
		util.ErrorInternalServer(c, "Failed to get paginated entities: "+err.Error())
		return
	}

	util.SuccessOK(c, "Paginated entities retrieved successfully", pagination)
}

func (ctrl *BaseController[T]) GetByField(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	if field == "" || value == "" {
		util.ErrorBadRequest(c, "Both 'field' and 'value' query parameters are required")
		return
	}

	entities, err := ctrl.Service.GetByField(field, value)
	if err != nil {
		util.ErrorInternalServer(c, "Failed to get entities by field: "+err.Error())
		return
	}

	util.SuccessOK(c, "Entities retrieved by field successfully", entities)
}

func (ctrl *BaseController[T]) FindByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		util.ErrorBadRequest(c, "Name query parameter is required")
		return
	}

	entities, err := ctrl.Service.FindByName(name)
	if err != nil {
		util.ErrorInternalServer(c, "Failed to find entities by name: "+err.Error())
		return
	}

	util.SuccessOK(c, "Entities found by name successfully", entities)
}