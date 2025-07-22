package util

import (
	"net/http"
	baseModel "user-management/model/base"

	"github.com/gin-gonic/gin"
)

// SuccessResponse creates a success response with appropriate status code
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := baseModel.BaseResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}

// ErrorResponse creates an error response with appropriate status code
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	response := baseModel.BaseResponse{
		Status:  statusCode,
		Message: message,
		Data:    nil,
	}
	c.JSON(statusCode, response)
}

// Helper functions for common responses
func SuccessOK(c *gin.Context, message string, data interface{}) {
	SuccessResponse(c, http.StatusOK, message, data)
}

func SuccessCreated(c *gin.Context, message string, data interface{}) {
	SuccessResponse(c, http.StatusCreated, message, data)
}

func ErrorBadRequest(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message)
}

func ErrorNotFound(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message)
}

func ErrorInternalServer(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message)
}

func ErrorUnauthorized(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message)
}

func ErrorForbidden(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message)
}