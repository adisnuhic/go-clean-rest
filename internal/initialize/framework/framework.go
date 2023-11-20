package framework

import (
	"github.com/adisnuhic/go-clean/config"
	"github.com/adisnuhic/go-clean/internal/initialize/framework/middlewares"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// InitGin returns GIN instance
func InitGin(l zerolog.Logger) *gin.Engine {
	g := gin.Default()

	// use middlewares
	g.Use(requestid.New())
	g.Use(middlewares.GinCorsMiddleware())
	g.Use(middlewares.GinZeroLogMiddleware(l))

	mode := gin.DebugMode
	environment := config.Load().Env
	if environment.Env == "production" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	return g
}
