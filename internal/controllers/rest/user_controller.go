package rest

import (
	"github.com/adisnuhic/go-clean/internal/services"
	"github.com/gin-gonic/gin"
)

// IUserController represents the user's controller/handler contract
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

// GetByID returns fake user data by ID
func (ctrl userController) GetByID(ctx *gin.Context) {
	user, errUser := ctrl.Service.GetByID(1)
	if errUser != nil {
		ctrl.RenderError(ctx, errUser)
	}

	ctrl.RenderSuccess(ctx, user)
}
