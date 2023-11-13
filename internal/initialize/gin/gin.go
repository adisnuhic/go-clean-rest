package gin

import (
	"github.com/adisnuhic/go-stanard-layout/config"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var zLogger zerolog.Logger

// returns instance of GIN framework
func Gin(l zerolog.Logger) *gin.Engine {

	g := gin.Default()
	zLogger = l

	// use request id
	g.Use(requestid.New())

	// use middlewares
	g.Use(corsMiddleware())
	g.Use(zeroLogMiddleware())

	// set debug mode
	mode := gin.DebugMode
	environment := config.Load().Env
	if environment.Env == "production" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	return g
}

// corsMiddleware -
func corsMiddleware() gin.HandlerFunc {
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

// zeroLogMiddleware middleware
func zeroLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", zLogger)
	}
}
