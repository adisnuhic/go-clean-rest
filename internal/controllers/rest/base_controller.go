package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/adisnuhic/go-clean/config"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/viewmodels"
	"github.com/adisnuhic/go-clean/pkg/apperror"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BaseController -
type BaseController struct {
}

// Render -
func (ctrl BaseController) Render(ctx *gin.Context, success bool, status int, data interface{}, err interface{}) {
	ctx.Status(status)

	response := viewmodels.ResponseBase{
		Success:   success,
		RequestID: requestid.Get(ctx),
		Data:      data,
		Error:     err,
	}

	if err != nil {
		switch err.(type) {
		case []apperror.ValidationError:
			response.Error = nil
			response.ValidationError = err
		}
	}

	ctx.JSON(status, response)
}

// RenderSuccess renders success response
func (ctrl BaseController) RenderSuccess(ctx *gin.Context, data interface{}) {
	ctrl.Render(ctx, true, http.StatusOK, data, nil)
}

// RenderError renders error response
func (ctrl BaseController) RenderError(ctx *gin.Context, err interface{}) {
	ctrl.Render(ctx, false, http.StatusInternalServerError, nil, err)
}

// RenderBadRequest renders bad request response
func (ctrl BaseController) RenderBadRequest(ctx *gin.Context, err interface{}) {
	ctrl.Render(ctx, false, http.StatusBadRequest, nil, err)
}

func (ctrl BaseController) RenderForbidden(ctx *gin.Context, err interface{}) {
	ctrl.Render(ctx, false, http.StatusForbidden, nil, err)
}

// RenderValidationError renders validation errors
func (ctrl BaseController) RenderValidationError(ctx *gin.Context, err error) {

	var vErr []apperror.ValidationError

	// if its model binding validation error
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range ve {
			myValidationErr := &apperror.ValidationError{}
			myValidationErr.Tag = fieldErr.Tag()
			myValidationErr.Field = fieldErr.Field()
			myValidationErr.Message = errTagToMessage(fieldErr)
			myValidationErr.Param = fieldErr.Param()
			vErr = append(vErr, *myValidationErr)
		}
	}

	// if its not model binding validation error it might be other errors (*json.UnmarshalTypeError, *json.SyntaxError, errors.errorString etc...)
	if _, ok := err.(validator.ValidationErrors); !ok {
		myValidationErr := &apperror.ValidationError{}
		myValidationErr.Message = err.Error()
		vErr = append(vErr, *myValidationErr)
	}

	ctrl.Render(ctx, false, http.StatusBadRequest, nil, vErr)
}

// Render data
func (ctrl BaseController) RenderData(ctx *gin.Context, status int, responseType string, data []byte, err interface{}) {
	ctx.Status(status)
	ctx.Data(status, responseType, data)
}

// RenderHTMLSuccess render success HTML data
func (ctrl BaseController) RenderHTMLSuccess(ctx *gin.Context, data []byte) {
	ctrl.RenderData(ctx, http.StatusOK, "text/html", data, nil)
}

// RenderCSVSuccess render success CSV data
func (ctrl BaseController) RenderCSVSuccess(ctx *gin.Context, data []byte) {
	ctrl.RenderData(ctx, http.StatusOK, "text/csv", data, nil)
}

// GetUserFromContext returns logged in user from context
func (ctrl BaseController) GetUserFromContext(ctx *gin.Context) (*models.User, *apperror.AppError) {
	user := &models.User{}
	bearer := ctx.Request.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		bearer = bearer[7:]
	}

	if bearer == "" {
		return nil, apperror.New(401, "unauthorized", "unauthorized")
	}

	if bearer != "" {
		token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Load().JWTConf.Secret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user.ID = uint64(claims["id"].(float64))
			user.Email = fmt.Sprintf("%v", claims["email"])
		}

		if err == nil && token.Valid {
			return user, nil
		}
	}

	return nil, apperror.New(401, "unauthorized", "unauthorized")
}

// errTagToMessage returns custom message for error tag
func errTagToMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "max":
		// if string
		if err.Kind().String() == "string" {
			return err.Field() + " length must not exceed " + err.Param() + " characters"
		} else {
			return err.Field() + " must be less than " + err.Param()
		}
	case "min":
		// if string
		if err.Kind().String() == "string" {
			return err.Field() + " length must be at least " + err.Param() + " characters"
		} else {
			return err.Field() + " must be greater than " + err.Param()
		}
	case "email":
		return err.Field() + " must be a valid email address"
	case "alphanum":
		return err.Field() + " must be alphanumeric only A-z, 0-9"
	case "number":
		return err.Field() + " must be a number only"
	case "iban":
		return err.Field() + " must be a valid IBAN"
	}

	return "An error occurred with field: " + err.Field()
}
