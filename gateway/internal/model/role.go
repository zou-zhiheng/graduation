package model

import (
	"time"
)

type Role struct {
	ID         uint64     `gorm:"column:id;type:bigint(20);primary_key" json:"id"`
	Name       string     `gorm:"column:name;type:varchar(20);NOT NULL" json:"name"`
	Code       string     `gorm:"column:code;type:varchar(10);NOT NULL" json:"code"`
	Api        string     `gorm:"column:api;type:text" json:"api"`
	Apis       []uint64   `gorm:"-"`
	RoleRoutes string     `gorm:"column:roleRoutes;type:varchar(20)" json:"roleroutes"`
	FirstPage  string     `gorm:"column:firstPage;type:varchar(20)" json:"firstpage"`
	Desc       string     `gorm:"column:desc;type:varchar(255)" json:"desc"`
	CreateTime *time.Time `gorm:"column:createTime;type:datetime;NOT NULL" json:"createtime"`
	UpdateTime *time.Time `gorm:"column:updateTime;type:datetime" json:"updatetime"`
}

// TableName table name
func (m *Role) TableName() string {
	return "role"
}
