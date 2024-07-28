package model

type CompanyUser struct {
	ID        uint64 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	CompanyID int64  `gorm:"column:company_id;type:bigint(20);NOT NULL" json:"companyId"`
	UserID    int64  `gorm:"column:user_id;type:bigint(20);NOT NULL" json:"userId"`
}

// TableName table name
func (m *CompanyUser) TableName() string {
	return "company_user"
}
