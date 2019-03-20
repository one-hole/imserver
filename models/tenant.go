package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// Tenant - model which will holds tenant's basic info
type Tenant struct {
	gorm.Model
	Name        string
	Password    string
	ReceivePort string         `gorm:"column:receive_port"`
	ServerCount int            `gorm:"default:0"`
	Enable      bool           `gorm:"default:true;index:idx_able_tenants"`
	Email       sql.NullString `gorm:"size:100;unique_index"`
	Phone       sql.NullString `gorm:"size:20;unique_index"`

	Servers []Server
}

// TableName set table's name
func (Tenant) TableName() string {
	return "tenants"
}
