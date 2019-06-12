package models

// Server belongs_to Tenant
type Server struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	TenantID uint   `gorm:"column:tenant_id;index"`
	Host     string `gorm:"column:host" json:"host"`
	Tag      string `gorm:"column:tag"`

	Tenant Tenant `gorm:"foreignkey:TenantID"`
}

// TableName set table's name
func (Server) TableName() string {
	return "servers"
}
