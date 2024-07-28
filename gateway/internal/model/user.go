package model

import (
	"time"
)

type User struct {
	ID         uint64     `gorm:"column:id;type:bigint(20);primary_key" json:"id"`
	Name       string     `gorm:"column:name;type:varchar(20);NOT NULL" json:"name"`
	Account    string     `gorm:"column:account;type:varchar(20);NOT NULL" json:"account"`
	Password   string     `gorm:"column:password;type:varchar(60);NOT NULL" json:"password"`
	AvatarUrl  string     `gorm:"column:avatarUrl;type:varchar(60);NOT NULL" json:"avatarurl"`
	IsValid    int        `gorm:"column:isValid;type:tinyint(1);default:1;NOT NULL" json:"isvalid"`
	Sex        string     `gorm:"column:sex;type:varchar(2);NOT NULL" json:"sex"`
	Phone      string     `gorm:"column:phone;type:varchar(11);NOT NULL" json:"phone"`
	Salt       string     `gorm:"column:salt;type:varchar(60);NOT NULL" json:"salt"`
	RoleId     string     `gorm:"column:roleId;type:text" json:"roleid"`
	RoleIds    []uint64   `gorm:"-"`
	CreateTime *time.Time `gorm:"column:createTime;type:datetime;NOT NULL" json:"createtime"`
	UpdateTime *time.Time `gorm:"column:updateTime;type:datetime" json:"updatetime"`
}

// TableName table name
func (m *User) TableName() string {
	return "user"
}
