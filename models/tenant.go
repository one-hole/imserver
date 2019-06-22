package models

import (
	"fmt"
	"github.com/one-hole/imserver/logs"
	"time"
)

// Tenant - model which will holds tenant's basic info
type Tenant struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	ExpireAt time.Time `gorm:"column:expire_time" json:"expire_at"`
	APIKey   string    `gorm:"column:api_key;unique_index" sql:"not null" json:"api_key"`
	Servers  []Server
}

// TableName set table's name
func (Tenant) TableName() string {
	return "tenants"
}

// Verify 校验链接是否有权限、目前的业务
// 1. 能找到商户
// 2. APIKey 没有过期
func (tenant *Tenant) Verify(key string) bool {
	
	logs.WebSocketLogger.Info(fmt.Sprintf("WebSocket: Tenant Key is: %s", key))
	
	if err := tenant.recordByKey(key); err != nil {
		return false
	}
	
	logs.WebSocketLogger.Info(fmt.Sprintf("WebSocket: Tenant Id is: %d", tenant.ID))
	return tenant.valid()
}

// recordByKey 查找商户
func (tenant *Tenant) recordByKey(key string) error {
	if result := db.Where(&Tenant{APIKey: key}).First(&tenant); result.Error != nil {
		return result.Error
	}
	return nil
}

// expired 判断 tenant 是否 valid?
func (tenant *Tenant) valid() bool {
	if tenant.ExpireAt.Unix() < time.Now().Unix() {
		return false
	}
	return true
}

// Hosts return's tenants servers
func (tenant *Tenant) Hosts() map[string]bool {
	hosts := make(map[string]bool)

	db.Model(&tenant).Related(&tenant.Servers)

	for _, server := range tenant.Servers {
		hosts[server.Host] = true
	}

	return hosts
}
