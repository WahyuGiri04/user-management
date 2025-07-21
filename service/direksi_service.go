package service

import (
	"errors"
	"user-management/model"
	"user-management/repository"
	baseService "user-management/service/base"
)

type DireksiServiceInterface interface {
	baseService.BaseServiceInterface[model.Direksi]
	// Add any specific methods for Direksi here if needed
	GetByCode(code string) (*model.Direksi, error)
}

type DireksiService struct {
	*baseService.BaseService[model.Direksi]
	repository repository.DireksiRepositoryInterface
}

func NewDireksiService(repo repository.DireksiRepositoryInterface) DireksiServiceInterface {
	return &DireksiService{
		BaseService: baseService.NewBaseService(repo),
		repository:  repo,
	}
}

// GetByCode - custom method specific to Direksi
func (s *DireksiService) GetByCode(code string) (*model.Direksi, error) {
	if code == "" {
		return nil, errors.New("code cannot be empty")
	}
	
	direksiList, err := s.repository.GetByField("code", code)
	if err != nil {
		return nil, err
	}
	
	if len(direksiList) == 0 {
		return nil, errors.New("direksi not found with the given code")
	}
	
	return &direksiList[0], nil
}