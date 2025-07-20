package baseModel

import "time"

type BaseModel struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UUID      string    `json:"uuid" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
}
