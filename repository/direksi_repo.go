package repository

import (
	"user-management/model"
	baseRepository "user-management/repository/base"
)

type DireksiRepositoryInterface interface {
	baseRepository.BaseRepositoryInterface[model.Direksi]
}

type DireksiRepository struct {
	*baseRepository.BaseRepository[model.Direksi]
}

func NewDireksiRepository() DireksiRepositoryInterface {
	return &DireksiRepository{
		BaseRepository: baseRepository.NewBaseRepository[model.Direksi](),
	}
}