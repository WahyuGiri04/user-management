package repository

import (
	"user-management/model"
	baseRepository "user-management/repository/base"
)

type RoleRepositoryInterface interface {
	baseRepository.BaseRepositoryInterface[model.Role]
}

type RoleRepository struct {
	*baseRepository.BaseRepository[model.Role]
}

func NewRoleRepository() RoleRepositoryInterface {
	return &RoleRepository{
		BaseRepository: baseRepository.NewBaseRepository[model.Role](),
	}
}