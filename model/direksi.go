package model

import (
	baseModel "user-management/model/base"
)

type Direksi struct {
	baseModel.BaseModel
	Nama string `json:"nama" gorm:"type:varchar(255)"`
}

func (Direksi) TableName() string {
	return "user-management.direksi"
}
