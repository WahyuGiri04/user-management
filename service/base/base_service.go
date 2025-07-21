package baseService

import (
	"errors"
	baseModel "user-management/model/base"
	baseRepository "user-management/repository/base"
)

type BaseServiceInterface[T any] interface {
	Create(entity *T) error
	GetAll(entities *[]T) error
	GetAllIncludingDeleted(entities *[]T) error
	GetByUUID(entity *T, uuid string) error
	Update(entity *T, uuid string) error
	Delete(uuid string) error
	SoftDelete(uuid string) error
	GetPagination(page, pageSize int, entities *[]T) (baseModel.Pagination, error)
	GetByField(field, value string) ([]T, error)
	FindByName(name string) ([]T, error)
}

type BaseService[T any] struct {
	Repository baseRepository.BaseRepositoryInterface[T]
}

func NewBaseService[T any](repo baseRepository.BaseRepositoryInterface[T]) *BaseService[T] {
	return &BaseService[T]{Repository: repo}
}

func (s *BaseService[T]) Create(entity *T) error {
	return s.Repository.Create(entity)
}

func (s *BaseService[T]) GetAll(entities *[]T) error {
	return s.Repository.GetAll(entities)
}

func (s *BaseService[T]) GetAllIncludingDeleted(entities *[]T) error {
	return s.Repository.GetAllIncludingDeleted(entities)
}

func (s *BaseService[T]) GetByUUID(entity *T, uuid string) error {
	if uuid == "" {
		return errors.New("UUID cannot be empty")
	}
	return s.Repository.GetByUUID(entity, uuid)
}

func (s *BaseService[T]) Update(entity *T, uuid string) error {
	if uuid == "" {
		return errors.New("UUID cannot be empty")
	}
	
	// Check if entity exists and is not soft deleted
	var existingEntity T
	if err := s.Repository.GetByUUID(&existingEntity, uuid); err != nil {
		return errors.New("entity not found or has been deleted")
	}
	
	return s.Repository.Update(entity, uuid)
}

// Delete - hard delete (permanent removal)
func (s *BaseService[T]) Delete(uuid string) error {
	if uuid == "" {
		return errors.New("UUID cannot be empty")
	}
	
	// Check if entity exists before deleting
	var existingEntity T
	if err := s.Repository.GetByUUID(&existingEntity, uuid); err != nil {
		return errors.New("entity not found or has been deleted")
	}
	
	return s.Repository.Delete(uuid)
}

// SoftDelete - marks entity as deleted without permanent removal
func (s *BaseService[T]) SoftDelete(uuid string) error {
	if uuid == "" {
		return errors.New("UUID cannot be empty")
	}
	
	// Check if entity exists and is not already soft deleted
	var existingEntity T
	if err := s.Repository.GetByUUID(&existingEntity, uuid); err != nil {
		return errors.New("entity not found or has already been deleted")
	}
	
	return s.Repository.SoftDelete(uuid)
}

func (s *BaseService[T]) GetPagination(page, pageSize int, entities *[]T) (baseModel.Pagination, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	return s.Repository.GetPagination(page, pageSize, entities)
}

func (s *BaseService[T]) GetByField(field, value string) ([]T, error) {
	if field == "" || value == "" {
		return nil, errors.New("field and value cannot be empty")
	}
	
	return s.Repository.GetByField(field, value)
}

func (s *BaseService[T]) FindByName(name string) ([]T, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	
	return s.Repository.FindByName(name)
}