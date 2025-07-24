package model

import baseModel "user-management/model/base"

type Role struct{
	baseModel.BaseModel
	Name string `json:"name" gorm:"type:varchar(255)"`
	Code string `json:"code" gorm:"type:varchar(255);unique"`
	Description string `json:"description" gorm:"type:varchar(500)"`
}

func (Role) TableName() string {
	return "user_management.role"
}