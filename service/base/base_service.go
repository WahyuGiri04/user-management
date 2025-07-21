package baseService

import (
	"errors"
	baseModel "user-management/model/base"
	baseRepository "user-management/repository/base"
)


type BaseServiceInterface[T any] interface {
	Create(entity *T) error
	GetAll(entities *[]T) error
	GetByUUID(entity *T, uuid string) error
	Update(entity *T, uuid string) error
	Delete(uuid string) error
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
	if err := s.Repository.Create(entity); err != nil {
		return err
	}
	return s.Repository.Create(entity)
}

func (s *BaseService[T]) GetAll(entities *[]T) error {
	if err := s.Repository.GetAll(entities); err != nil {
		return err
	}
	return nil
}

func (s *BaseService[T]) GetByUUID(entity *T, uuid string) error {
	if err := s.Repository.GetByUUID(entity, uuid); err != nil {
		return err
	}
	return nil
}

func (s *BaseService[T]) Update(entity *T, uuid string) error {
	var existingEntity T
	if err := s.Repository.GetByUUID(&existingEntity, uuid); err != nil {
		return err
	}
	return s.Repository.Update(entity, uuid)
}

func (s *BaseService[T]) Delete(uuid string) error {
	return s.Repository.Delete(uuid)
}

func (s *BaseService[T]) GetPagination(page, pageSize int, entities *[]T) (baseModel.Pagination, error) {
	if page < 0 {
		page = 1
	}

	if pageSize < 0 || pageSize > 100 {
		pageSize = 10
	}

	return s.Repository.GetPagination(page, pageSize, entities)
}

func (s *BaseService[T]) GetByField(field, value string) ([]T, error) {
	// Business logic validation
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