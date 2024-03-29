package rest

import (
	"github.com/adisnuhic/go-clean/internal/requests"
	"github.com/adisnuhic/go-clean/internal/services"
	"github.com/adisnuhic/go-clean/internal/viewmodels"
	"github.com/gin-gonic/gin"
)

// IAccountController represents the account controller contract
type IAccountController interface {
	Login(ctx *gin.Context)
}

type accountController struct {
	BaseController
	Service services.IAccountService
}

// NewAccountController -
func NewAccountController(svc services.IAccountService) IAccountController {
	return &accountController{
		Service: svc,
	}
}

// Ping returns pong
func (ctrl accountController) Login(ctx *gin.Context) {
	reqObj := requests.Login{}

	if err := ctx.ShouldBindJSON(&reqObj); err != nil {
		ctrl.RenderValidationError(ctx, err)
		ctx.Abort()
		return
	}

	user, accessToken, refreshToken, appErr := ctrl.Service.Login(reqObj.Email, reqObj.Password)

	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		ctx.Abort()
		return
	}

	ctrl.RenderSuccess(ctx, &viewmodels.Auth{
		User:         user,
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}
