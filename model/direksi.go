package model

import (
	baseModel "user-management/model/base"
)

type Direksi struct {
	baseModel.BaseModel
	Nama string `json:"nama" gorm:"type:varchar(255)"`
	Code string `json:"code" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(500)"`
}

func (Direksi) TableName() string {
	return "user-management.direksi"
}
