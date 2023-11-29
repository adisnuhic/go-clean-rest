package rest

import (
	"github.com/adisnuhic/go-clean/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	user, errUser := ctrl.Service.GetByID(12)

	if errUser != nil && gorm.IsRecordNotFoundError(errUser) {
		ctrl.RenderNotFound(ctx, errUser)
		return
	}

	if errUser != nil {
		ctrl.RenderError(ctx, errUser)
		return
	}

	ctrl.RenderSuccess(ctx, user)
}
