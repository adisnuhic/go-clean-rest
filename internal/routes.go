package internal

import (
	"github.com/adisnuhic/go-clean/internal/controllers/rest"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/pkg/container"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	userCtrl    rest.IUserController
	healthCtrl  rest.IHealthController
	accountCtrl rest.IAccountController
)

// initialize app routes
func InitRoutes(c container.Container, app *gin.Engine) {

	// Resolve dependencies and return concrete type of given abstractions
	c.Make(&userCtrl)

	// --------------------------------------------------- //
	//		        		  V1 API			       	   //
	// --------------------------------------------------- //
	// Group routes to specific version
	v1 := app.Group("/v1")

	// health check route
	healthRoutes := v1.Group("/ping")
	healthRoutes.GET("/", healthCtrl.Ping)

	// account routes
	accountRoutes := v1.Group("/account")
	accountRoutes.POST("/login", accountCtrl.Login)

	// --------------------------------------------------- //
	//						SWAGGER						   //
	// --------------------------------------------------- //
	// Swagger API documentation
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
