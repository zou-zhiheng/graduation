package model

import (
	"time"
)

type Api struct {
	ID         uint64     `gorm:"column:id;type:bigint(20);primary_key" json:"id"`
	Name       string     `gorm:"column:name;type:varchar(20);NOT NULL" json:"name"`
	Url        string     `gorm:"column:url;type:varchar(20);NOT NULL" json:"url"`
	Method     string     `gorm:"column:method;type:varchar(255);NOT NULL" json:"method"`
	Desc       string     `gorm:"column:desc;type:varchar(255)" json:"desc"`
	CreateTime *time.Time `gorm:"column:createTime;type:datetime;NOT NULL" json:"createtime"`
	UpdateTime *time.Time `gorm:"column:updateTime;type:datetime" json:"updatetime"`
}

// TableName table name
func (m *Api) TableName() string {
	return "api"
}
