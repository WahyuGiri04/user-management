package baseModel

import "time"

type BaseModel struct {
	ID        int       `json:"-" gorm:"primaryKey"`
	UUID      string    `json:"uuid" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(255)"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(255)"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
}
