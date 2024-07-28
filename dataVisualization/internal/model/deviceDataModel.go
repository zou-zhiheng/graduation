package model

import (
	"gorm.io/gorm"
	"time"
)

type DeviceDataModel struct {
	ID         uint64     `gorm:"column:id;type:bigint(20);primary_key" json:"id"`
	Data       string     `gorm:"column:data;type:text" json:"data"`
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"createTime"`
}

type DataDetail struct {
	Key   string
	Value float32
	Unit  string
}

// TableName table name
func (m *DeviceDataModel) TableName() string {
	return "device_data_model"
}

// NewDeviceDataModel NewUser 创建一个函数来接受表名并返回对应的模型结构体
func NewDeviceDataModel(tableName string) interface{} {
	return struct {
		gorm.Model
		Name  string
		Email string
	}{Model: gorm.Model{}, Name: tableName}
}
