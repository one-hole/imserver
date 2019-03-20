package models

import (
	"github.com/jinzhu/gorm"
)

// Tenant - model which will holds tenant's basic info
type Tenant struct {
	gorm.Model
	Name        string
	Email       string `gorm:"size:100;unique_index"`
	Phone       string `gorm:"size:20;unique_index"`
	Password    string
	ReceivePort string `gorm:"column:receive_port"`
	ServerCount int    `gorm:"default:0"`
	Enable      bool   `gorm:"default:true;index:idx_able_tenants"`

	Servers []Server
}

// TableName set table's name
func (Tenant) TableName() string {
	return "tenants"
}
