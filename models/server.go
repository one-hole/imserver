package models

import (
	"github.com/jinzhu/gorm"
)

// Server belongs_to Tenant
type Server struct {
	gorm.Model
	TenantID uint   `gorm:"column:tenant_id;index"`
	Tenant   Tenant `gorm:"foreignkey:TenantID"`
	Enable   bool   `gorm:"default:true"`
	Address  string
}

// TableName set table's name
func (Server) TableName() string {
	return "servers"
}
