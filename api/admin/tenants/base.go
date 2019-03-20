package tenants

import "github.com/w-zengtao/socket-server/models"

func loadTenant(id uint) *models.Tenant {
	tenant := &models.Tenant{}
	models.DB.First(&tenant, id)
	return tenant
}
