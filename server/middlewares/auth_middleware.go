package middlewares

import (
	"github.com/gin-gonic/gin"
	services "github.com/lucasdk3/maui-oqcomer-api/services"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_schema):]

		if !services.JWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
