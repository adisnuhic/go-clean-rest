package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
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
func GinZeroLogMiddleware(l zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", l)
	}
}
