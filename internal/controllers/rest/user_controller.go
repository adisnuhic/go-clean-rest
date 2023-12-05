package rest

import (
	"strconv"

	"github.com/adisnuhic/go-clean/internal/ecode"
	"github.com/adisnuhic/go-clean/internal/services"
	"github.com/adisnuhic/go-clean/pkg/apperror"
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
	userID, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
	if err != nil {
		ctrl.RenderBadRequest(ctx, apperror.New(ecode.ErrRequestParamValidationCode, err.Error(), ecode.ErrRequestParamValidationMsg))
		return
	}

	user, appErr := ctrl.Service.GetByID(userID)

	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		ctx.Abort()
		return
	}

	ctrl.RenderSuccess(ctx, user)
}
