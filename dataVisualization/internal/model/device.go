package model

import "time"

type Device struct {
	ID           uint64     `gorm:"column:id;type:bigint(20);primary_key" json:"id"`
	Name         string     `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Code         string     `gorm:"column:code;type:varchar(255);NOT NULL" json:"code"`
	IsCustom     int        `gorm:"column:is_custom;type:smallint;NOT NULL" json:"isCustom"`
	Protocol     string     `gorm:"column:protocol;type:text;NOT NULL" json:"protocol"`
	DeviceTypeID uint64     `gorm:"column:device_type_id;type:bigint(20);NOT NULL" json:"deviceTypeId"`
	CreateTime   *time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"createtime"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"updatetime"`
}

// TableName table name
func (m *Device) TableName() string {
	return "device"
}
