package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ALOWED_ORIGINS = "http://localhost:3000,https://harmony.businessbuilders.city,https://api.harmony.businessbuilders.city,https://data.harmony.businessbuilders.city"
)

func IsOriginAllowed(origin string) bool {
	allowedOrigins := strings.Split(ALOWED_ORIGINS, ",")

	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}

	return false
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		origin := c.Request.Header.Get("Origin")

		if IsOriginAllowed(origin) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, harmony_micro_services")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
