package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int    `gorm:"column:id; type:string; not null; unique; primaryKey; autoIncrement; comment:'primary Key is ID'"`
	RoleName string `gorm:"column:role_name"`
	IsActive bool   `gorm:"column:is_active; type:boolean"`
	RoleNote string `gorm:"column:role_note; type:text"`
}

func (r *Role) TableName() string {
	return "db_role"
}
