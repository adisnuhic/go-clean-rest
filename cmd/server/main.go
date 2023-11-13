package server

import (
	"github.com/adisnuhic/go-stanard-layout/config"
	"github.com/adisnuhic/go-stanard-layout/internal/initialize/db"
	"github.com/adisnuhic/go-stanard-layout/internal/initialize/gin"
	"github.com/adisnuhic/go-stanard-layout/internal/initialize/log"
)

func main() {

	// load app config
	cfg := config.Load()

	// init logger
	logger := log.NewLogger()

	// initialize database
	db.Init(cfg)

	// initialize GIN framework
	app := gin.Gin(logger)

	// Run app
	app.Run(":8080")
}
