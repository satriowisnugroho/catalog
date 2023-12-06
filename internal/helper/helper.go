package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/catalog/internal/entity/types"
)

func StringInArray(target string, arr []string) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}

func GetTenant(c *gin.Context) types.TenantType {
	return types.TenantTypeNameToValue[c.GetHeader("X-Tenant")]
}
