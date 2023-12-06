package config

const (
	// ServiceName is the name of this service
	ServiceName = "catalog"
	// UniqueConstraintViolationCode is the pgError code for unique constraint violation error
	UniqueConstraintViolationCode = "23505"
	// SKUTenantUniqueConstraint is the name sku and tenant index name
	SKUTenantUniqueConstraint = "products_sku_tenant_idx"
)
