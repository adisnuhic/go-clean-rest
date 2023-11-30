package rest

import (
	"github.com/adisnuhic/go-clean/internal/services"
	"github.com/gin-gonic/gin"
)

// IUserController represents the user controller contract
type IUserController interface {
	GetByID(ctx *gin.Context)
}

type userController struct {
	BaseController
	Service services.IUserService
}

// NewUserController -
func NewUserController(svc services.IUserService) IUserController {
	return &userController{
		Service: svc,
	}
}

// GetByID returns user by ID
func (ctrl userController) GetByID(ctx *gin.Context) {
	user, appErr := ctrl.Service.GetByID(12)

	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		ctx.Abort()
		return
	}

	ctrl.RenderSuccess(ctx, user)
}
