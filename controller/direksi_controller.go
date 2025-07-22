package controller

import (
	"user-management/model"
	"user-management/service"
	"user-management/repository"
	"user-management/util"
	baseController "user-management/controller/base"

	"github.com/gin-gonic/gin"
)

type DireksiControllerInterface interface {
	baseController.BaseControllerInterface[model.Direksi]
	// Add any specific methods for Direksi here if needed
	GetByCode(c *gin.Context)
}

type DireksiController struct {
	*baseController.BaseController[model.Direksi]
	service service.DireksiServiceInterface
}

func NewDireksiController() DireksiControllerInterface {
	direksiRepo := repository.NewDireksiRepository()
	direksiService := service.NewDireksiService(direksiRepo)
	
	return &DireksiController{
		BaseController: baseController.NewBaseController(direksiService),
		service:        direksiService,
	}
}

// GetByCode - custom method specific to Direksi
func (ctrl *DireksiController) GetByCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		util.ErrorBadRequest(c, "Code query parameter is required")
		return
	}

	direksi, err := ctrl.service.GetByCode(code)
	if err != nil {
		util.ErrorNotFound(c, "Direksi not found: "+err.Error())
		return
	}

	util.SuccessOK(c, "Direksi retrieved by code successfully", direksi)
}