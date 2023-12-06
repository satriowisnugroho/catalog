package middleware

import (
	"github.com/gin-gonic/gin"
)

func Tenant() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Set tenant

		c.Next()
	}
}
