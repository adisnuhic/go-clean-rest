package main

import (
	"github.com/adisnuhic/go-clean/config"
	"github.com/adisnuhic/go-clean/internal"
	"github.com/adisnuhic/go-clean/internal/controllers/rest"
	"github.com/adisnuhic/go-clean/internal/initialize/framework"
	"github.com/adisnuhic/go-clean/internal/initialize/log"
	"github.com/adisnuhic/go-clean/internal/initialize/mysql"
	repositories "github.com/adisnuhic/go-clean/internal/repositories/mysql"
	"github.com/adisnuhic/go-clean/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container"
)

var app *gin.Engine

func main() {

	// load app config
	cfg := config.Load()

	// init logger
	logger := log.NewLogger()

	// initialize mysql database
	mysql.Init(cfg)

	// create new DI container
	c := container.NewContainer()

	// init repositories
	repositories.Init(c)

	// init services/usecases
	services.Init(c)

	// init controllers/handlers
	rest.Init(c)

	// initialize framework
	app = framework.InitGin(logger)

	// initialize routes
	internal.InitRoutes(c, app)

	// Run the app
	app.Run(":8080")
}
