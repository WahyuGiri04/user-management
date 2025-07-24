package controller

import (
	baseController "user-management/controller/base"
	"user-management/model"
	"user-management/repository"
	"user-management/service"
)

type RoleControllerInterface interface{
	baseController.BaseControllerInterface[model.Role]
}

type RoleController struct {
	*baseController.BaseController[model.Role]
	service service.RoleServiceInterface
}

func NewRoleController() RoleControllerInterface {
	roleRepo := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepo)

	return &RoleController{
		BaseController: baseController.NewBaseController(roleService),
		service: roleService,
	}
}

