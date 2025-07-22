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

// Create - override base Create method with code validation
func (s *DireksiService) Create(entity *model.Direksi) error {
	// Validate that code is not empty
	if entity.Code == "" {
		return errors.New("code cannot be empty")
	}
	
	// Check if code already exists
	existingDireksi, err := s.GetByCode(entity.Code)
	if err == nil && existingDireksi != nil {
		return errors.New("direksi with this code already exists")
	}
	
	// If error is not "not found", then it's a real error
	if err != nil && err.Error() != "direksi not found with the given code" {
		return err
	}
	
	// Code is unique, proceed with creation
	return s.repository.Create(entity)
}

// Update - override base Update method with code validation
func (s *DireksiService) Update(entity *model.Direksi, uuid string) error {
	if uuid == "" {
		return errors.New("UUID cannot be empty")
	}
	
	if entity.Code == "" {
		return errors.New("code cannot be empty")
	}
	
	// Check if entity exists and is not soft deleted
	var existingEntity model.Direksi
	if err := s.repository.GetByUUID(&existingEntity, uuid); err != nil {
		return errors.New("entity not found or has been deleted")
	}
	
	// Check if the new code is different from the existing one
	if existingEntity.Code != entity.Code {
		// Check if the new code already exists for another entity
		existingDireksiWithCode, err := s.GetByCode(entity.Code)
		if err == nil && existingDireksiWithCode != nil && existingDireksiWithCode.UUID.String() != uuid {
			return errors.New("direksi with this code already exists")
		}
		
		// If error is not "not found", then it's a real error
		if err != nil && err.Error() != "direksi not found with the given code" {
			return err
		}
	}
	
	// Code validation passed, proceed with update
	return s.repository.Update(entity, uuid)
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