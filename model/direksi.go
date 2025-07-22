package model

import (
	baseModel "user-management/model/base"
)

type Direksi struct {
	baseModel.BaseModel
	Name string `json:"name" gorm:"type:varchar(255)"`
	Code string `json:"code" gorm:"type:varchar(255);unique"`
	Description string `json:"description" gorm:"type:varchar(500)"`
}

func (Direksi) TableName() string {
	return "user_management.direksi"
}
