package model

import (
	"time"
)

type UserDevice struct {
	ID         uint64     `gorm:"column:id;type:bigint(20);primary_key" json:"id"`
	UserID     uint64     `gorm:"column:user_id;type:bigint(20);NOT NULL" json:"userId"`
	DeviceID   uint64     `gorm:"column:device_id;type:bigint(20);NOT NULL" json:"deviceId"`
	UserName   string     `gorm:"column:user_name;type:varchar(255);NOT NULL" json:"userName"`
	DeviceName string     `gorm:"column:device_name;type:varchar(255);NOT NULL" json:"deviceName"`
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"createtime"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime" json:"updatetime"`
	DeviceCode string     `gorm:"column:device_code;type:varchar(255);NOT NULL" json:"deviceCode"`
}

// TableName table name
func (m *UserDevice) TableName() string {
	return "user_device"
}
