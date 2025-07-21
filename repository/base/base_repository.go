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
	GetByUUID(entity *T, uuid string) error
	Update(entity *T, uuid string) error
	Delete(uuid string) error
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

func (r *BaseRepository[T]) GetAll(entities *[]T) error {
	return r.DB.Find(entities).Error
}

func (r *BaseRepository[T]) GetByUUID(entity *T, uuid string) error {
	// return r.DB.Where("uuid = ?", uuid).First(entity).Error
	return r.DB.First(entity, uuid).Error
}

func (r *BaseRepository[T]) Update(entity *T, uuid string) error {
	return r.DB.Model(entity).Where("uuid = ?", uuid).Updates(entity).Error
}

func (r *BaseRepository[T]) Delete(uuid string) error {
	return r.DB.Delete(new(T), uuid).Error
}

func (r *BaseRepository[T]) GetPagination(page, pageSize int, entities *[]T) (baseModel.Pagination, error) {
	var totalRows int64
	r.DB.Model(new(T)).Count(&totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	offset := (page - 1) * pageSize

	err := r.DB.Limit(pageSize).Offset(offset).Find(entities).Error
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

func (r *BaseRepository[T]) GetByField(field, value string) ([]T, error) {
	var entities []T
	err := r.DB.Where(field+" LIKE ?", "%"+value+"%").Find(&entities).Error
	return entities, err
}

func (r *BaseRepository[T]) FindByName(name string) ([]T, error) {
	var entities []T
	err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&entities).Error
	return entities, err
}