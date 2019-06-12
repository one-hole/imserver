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

// RecordByHost 通过 Host 查找 Server
func (server *Server) RecordByHost(host string) error {
	if result := DB.Where(&Server{Host: host}).First(&server); result.Error != nil {
		return result.Error
	}
	return nil
}

// Valid true valid & false invalid
func (server *Server) Valid() bool {
	DB.Model(&server).Related(&server.Tenant)
	return server.Tenant.valid()
}
