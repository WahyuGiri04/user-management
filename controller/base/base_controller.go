package baseController

import (
	"net/http"
	"strconv"
	baseModel "user-management/model/base"
	baseService "user-management/service/base"

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

// Helper function to create success response
func (ctrl *BaseController[T]) SuccessResponse(c *gin.Context, message string, data interface{}) {
	response := baseModel.BaseResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

// Helper function to create error response
func (ctrl *BaseController[T]) ErrorResponse(c *gin.Context, statusCode int, message string) {
	response := baseModel.BaseResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	}
	c.JSON(statusCode, response)
}

func (ctrl *BaseController[T]) Create(c *gin.Context) {
	var entity T
	if err := c.ShouldBindJSON(&entity); err != nil {
		ctrl.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	if err := ctrl.Service.Create(&entity); err != nil {
		ctrl.ErrorResponse(c, http.StatusInternalServerError, "Failed to create entity: "+err.Error())
		return
	}

	ctrl.SuccessResponse(c, "Entity created successfully", entity)
}

func (ctrl *BaseController[T]) GetAll(c *gin.Context) {
	var entities []T
	if err := ctrl.Service.GetAll(&entities); err != nil {
		ctrl.ErrorResponse(c, http.StatusInternalServerError, "Failed to get entities: "+err.Error())
		return
	}

	ctrl.SuccessResponse(c, "Entities retrieved successfully", entities)
}

func (ctrl *BaseController[T]) GetByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		ctrl.ErrorResponse(c, http.StatusBadRequest, "UUID parameter is required")
		return
	}

	var entity T
	if err := ctrl.Service.GetByUUID(&entity, uuid); err != nil {
		ctrl.ErrorResponse(c, http.StatusNotFound, "Entity not found: "+err.Error())
		return
	}

	ctrl.SuccessResponse(c, "Entity retrieved successfully", entity)
}

func (ctrl *BaseController[T]) Update(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		ctrl.ErrorResponse(c, http.StatusBadRequest, "UUID parameter is required")
		return
	}

	var entity T
	if err := c.ShouldBindJSON(&entity); err != nil {
		ctrl.ErrorResponse(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	if err := ctrl.Service.Update(&entity, uuid); err != nil {
		ctrl.ErrorResponse(c, http.StatusInternalServerError, "Failed to update entity: "+err.Error())
		return
	}

	ctrl.SuccessResponse(c, "Entity updated successfully", entity)
}

func (ctrl *BaseController[T]) Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		ctrl.ErrorResponse(c, http.StatusBadRequest, "UUID parameter is required")
		return
	}

	if err := ctrl.Service.Delete(uuid); err != nil {
		ctrl.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete entity: "+err.Error())
		return
	}

	ctrl.SuccessResponse(c, "Entity deleted successfully", nil)
}

func (ctrl *BaseController[T]) SoftDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		ctrl.ErrorResponse(c, http.StatusBadRequest, "UUID parameter is required")
		return
	}

	if err := ctrl.Service.SoftDelete(uuid); err != nil {
		ctrl.ErrorResponse(c, http.StatusInternalServerError, "Failed to soft delete entity: "+err.Error())
		return
	}

	ctrl.SuccessResponse(c, "Entity soft deleted successfully", nil)
}

func (ctrl *BaseController[T]) GetPagination(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctrl.ErrorResponse(c, http.StatusBadRequest, "Invalid page parameter")
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		ctrl.ErrorResponse(c, http.StatusBadRequest, "Invalid page_size parameter")
		return
	}

	var entities []T
	pagination, err := ctrl.Service.GetPagination(page, pageSize, &entities)
	if err != nil {
		ctrl.ErrorResponse(c, http.StatusInternalServerError, "Failed to get paginated entities: "+err.Error())
		return
	}

	ctrl.SuccessResponse(c, "Paginated entities retrieved successfully", pagination)
}

func (ctrl *BaseController[T]) GetByField(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	if field == "" || value == "" {
		ctrl.ErrorResponse(c, http.StatusBadRequest, "Both 'field' and 'value' query parameters are required")
		return
	}

	entities, err := ctrl.Service.GetByField(field, value)
	if err != nil {
		ctrl.ErrorResponse(c, http.StatusInternalServerError, "Failed to get entities by field: "+err.Error())
		return
	}

	ctrl.SuccessResponse(c, "Entities retrieved by field successfully", entities)
}

func (ctrl *BaseController[T]) FindByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		ctrl.ErrorResponse(c, http.StatusBadRequest, "Name query parameter is required")
		return
	}

	entities, err := ctrl.Service.FindByName(name)
	if err != nil {
		ctrl.ErrorResponse(c, http.StatusInternalServerError, "Failed to find entities by name: "+err.Error())
		return
	}

	ctrl.SuccessResponse(c, "Entities found by name successfully", entities)
}