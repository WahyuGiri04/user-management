package service

import (
	"user-management/model"
	"user-management/repository"
	baseService "user-management/service/base"
)

type RoleServiceInterface interface {
	baseService.BaseServiceInterface[model.Role]
}

type RoleService struct {
	*baseService.BaseService[model.Role]
	repository repository.RoleRepositoryInterface
}

func NewRoleService(repo repository.RoleRepositoryInterface) RoleServiceInterface {
	return &RoleService{
		BaseService: baseService.NewBaseService(repo),
		repository: repo,
	}
}