package model

import (
	"time"
)

type Role struct {
	ID         uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name       string     `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Code       string     `gorm:"column:code;type:varchar(255);NOT NULL" json:"code"`
	Api        string     `gorm:"column:api;type:text" json:"api"`
	RoleRoutes string     `gorm:"column:roleRoutes;type:varchar(255)" json:"roleroutes"`
	FirstPage  string     `gorm:"column:firstPage;type:varchar(255)" json:"firstpage"`
	Desc       string     `gorm:"column:desc;type:varchar(255)" json:"desc"`
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"createtime"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime" json:"updatetime"`
}

// TableName table name
func (m *Role) TableName() string {
	return "role"
}
