package rest

import (
	"github.com/gin-gonic/gin"
)

// IHealthController represents the user's controller contract
type IHealthController interface {
	Ping(ctx *gin.Context)
}

type healthController struct {
	BaseController
}

// NewHealthController -
func NewHealthController() IHealthController {
	return &healthController{}
}

// Ping returns pong
func (ctrl healthController) Ping(ctx *gin.Context) {
	ctrl.RenderSuccess(ctx, "PONG")
}
