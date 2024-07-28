package model

import "time"

type DeviceType struct {
	ID         uint64     `gorm:"column:id;type:bigint(20);primary_key" json:"id"`
	Name       string     `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Protocol   string     `gorm:"column:protocol;type:text;NOT NULL" json:"protocol"`
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"createtime"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"updatetime"`
}

// TableName table name
func (m *DeviceType) TableName() string {
	return "device_type"
}
