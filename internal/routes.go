package internal

import (
	"github.com/adisnuhic/go-clean/internal/controllers/rest"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/pkg/container"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	userCtrl rest.IUserController
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
	healthRoutes := v1.Group("/users")
	healthRoutes.GET("/", userCtrl.GetByID)

	// --------------------------------------------------- //
	//						SWAGGER						   //
	// --------------------------------------------------- //
	// Swagger API documentation
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
