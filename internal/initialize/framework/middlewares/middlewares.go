package middlewares

import (
	"github.com/adisnuhic/go-clean/pkg/log"
	"github.com/gin-gonic/gin"
)

// GinCorsMiddleware return CORS middleware for GIN framework
func GinCorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// GinZeroLogMiddleware return Zerolog middleware for GIN framework
func GinZeroLogMiddleware(l log.ILogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", l)
	}
}
