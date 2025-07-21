package baseRepository

import (
	"user-management/config"
	baseModel "user-management/model/base"

	"math"

	"gorm.io/gorm"
)

type BaseRepositoryInterface[T any] interface {
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
	GetDB() *gorm.DB
}

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	return &BaseRepository[T]{DB: config.DB}
}

func (r *BaseRepository[T]) GetDB() *gorm.DB {
	return r.DB
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

// GetAll - only returns active (not soft deleted) entities
func (r *BaseRepository[T]) GetAll(entities *[]T) error {
	return r.DB.Where("is_deleted = ? AND is_active = ?", false, true).Find(entities).Error
}

// GetAllIncludingDeleted - returns all entities including soft deleted ones
func (r *BaseRepository[T]) GetAllIncludingDeleted(entities *[]T) error {
	return r.DB.Unscoped().Find(entities).Error
}

// GetByUUID - only returns active (not soft deleted) entity
func (r *BaseRepository[T]) GetByUUID(entity *T, uuid string) error {
	return r.DB.Where("uuid = ? AND is_deleted = ? AND is_active = ?", uuid, false, true).First(entity).Error
}

func (r *BaseRepository[T]) Update(entity *T, uuid string) error {
	return r.DB.Model(entity).Where("uuid = ? AND is_deleted = ?", uuid, false).Updates(entity).Error
}

// Delete - hard delete (permanent removal)
func (r *BaseRepository[T]) Delete(uuid string) error {
	return r.DB.Where("uuid = ?", uuid).Delete(new(T)).Error
}

// SoftDelete - marks entity as deleted by setting is_deleted = true and is_active = false
func (r *BaseRepository[T]) SoftDelete(uuid string) error {
	return r.DB.Model(new(T)).Where("uuid = ?", uuid).Updates(map[string]interface{}{
		"is_deleted": true,
		"is_active":  false,
	}).Error
}

// GetPagination - only returns active (not soft deleted) entities with pagination
func (r *BaseRepository[T]) GetPagination(page, pageSize int, entities *[]T) (baseModel.Pagination, error) {
	var totalRows int64
	r.DB.Model(new(T)).Where("is_deleted = ? AND is_active = ?", false, true).Count(&totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	offset := (page - 1) * pageSize

	err := r.DB.Where("is_deleted = ? AND is_active = ?", false, true).
		Limit(pageSize).Offset(offset).Find(entities).Error
	if err != nil {
		return baseModel.Pagination{}, err
	}

	return baseModel.Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       entities,
	}, nil
}

// GetByField - only returns active (not soft deleted) entities
func (r *BaseRepository[T]) GetByField(field, value string) ([]T, error) {
	var entities []T
	err := r.DB.Where(field+" LIKE ? AND is_deleted = ? AND is_active = ?", "%"+value+"%", false, true).
		Find(&entities).Error
	return entities, err
}

// FindByName - only returns active (not soft deleted) entities
func (r *BaseRepository[T]) FindByName(name string) ([]T, error) {
	var entities []T
	err := r.DB.Where("nama LIKE ? AND is_deleted = ? AND is_active = ?", "%"+name+"%", false, true).
		Find(&entities).Error
	return entities, err
}