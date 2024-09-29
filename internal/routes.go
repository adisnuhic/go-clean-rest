package internal

import (
	"github.com/adisnuhic/go-clean/internal/controllers/rest"
	middleware "github.com/adisnuhic/go-clean/internal/middlewares"
	"github.com/adisnuhic/go-clean/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/pkg/container"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	healthCtrl  rest.IHealthController
	userCtrl    rest.IUserController
	accountCtrl rest.IAccountController
)

// initialize app routes
func InitRoutes(c container.Container, app *gin.Engine, logger log.ILogger) {

	// Resolve dependencies and return concrete type of given abstractions
	c.Make(&healthCtrl)
	c.Make(&userCtrl)
	c.Make(&accountCtrl)

	// --------------------------------------------------- //
	//		        		  V1 API			       	   //
	// --------------------------------------------------- //
	// Group routes to specific version
	v1 := app.Group("/v1")

	// auth routes
	authRoutes := v1.Group("auth", middleware.Authorization(logger))
	authRoutes.GET("/", healthCtrl.Ping)

	// health check route
	healthRoutes := v1.Group("/ping")
	healthRoutes.GET("/", healthCtrl.Ping)

	// account routes
	accountRoutes := v1.Group("/account")
	accountRoutes.POST("/login", accountCtrl.Login)

	// account routes
	usersRoutes := v1.Group("/users")
	usersRoutes.GET("/:id", middleware.ProfileMiddleware(userCtrl.GetByID, logger))

	// --------------------------------------------------- //
	//						SWAGGER						   //
	// --------------------------------------------------- //
	// Swagger API documentation
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
