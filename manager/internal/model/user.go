package model

import (
	"time"
)

type User struct {
	ID         uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name       string     `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Account    string     `gorm:"column:account;type:varchar(255);NOT NULL" json:"account"`
	Password   string     `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	AvatarUrl  string     `gorm:"column:avatar_url;type:varchar(255);NOT NULL" json:"avatarurl"`
	IsValid    int        `gorm:"column:is_Valid;type:tinyint(1);default:1;NOT NULL" json:"isvalid"`
	Sex        string     `gorm:"column:sex;type:varchar(2);NOT NULL" json:"sex"`
	Phone      string     `gorm:"column:phone;type:varchar(255);NOT NULL" json:"phone"`
	Salt       string     `gorm:"column:salt;type:varchar(255);NOT NULL" json:"salt"`
	RoleId     string     `gorm:"column:role_id;type:text" json:"roleid"`
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"createtime"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"updatetime"`
}

// TableName table name
func (m *User) TableName() string {
	return "user"
}
