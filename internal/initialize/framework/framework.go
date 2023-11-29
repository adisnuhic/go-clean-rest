package framework

import (
	"github.com/adisnuhic/go-clean/config"
	"github.com/adisnuhic/go-clean/internal/initialize/framework/middlewares"
	"github.com/adisnuhic/go-clean/pkg/log"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// InitGin returns GIN instance
func InitGin(logger log.ILogger) *gin.Engine {
	g := gin.Default()

	// use middlewares
	g.Use(requestid.New())
	g.Use(middlewares.GinCorsMiddleware())
	g.Use(middlewares.GinZeroLogMiddleware(logger))

	mode := gin.DebugMode
	environment := config.Load().Env
	if environment.Env == "production" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	return g
}
