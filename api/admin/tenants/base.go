package tenants

import "github.com/one-hole/imserver/models"

func loadTenant(id uint) *models.Tenant {
	tenant := &models.Tenant{}
	models.DB.First(&tenant, id)
	return tenant
}
