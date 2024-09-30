package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/adisnuhic/go-clean/config"
	"github.com/adisnuhic/go-clean/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Authorization middleware
func Authorization(logger log.ILogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := findAuthorizationToken(ctx.Request)

		// Validate token
		if tokenStr != "" {
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.Load().JWTConf.Secret), nil
			})

			if err != nil || !token.Valid {
				errAbort := ctx.AbortWithError(401, errors.New("unauthorized"))
				if errAbort != nil {
					logger.Errorf("Abort with error unauthorized err: %v", errAbort)
				}
				return
			}
		}

		if tokenStr == "" {
			errAbort := ctx.AbortWithError(401, errors.New("unauthorized"))
			if errAbort != nil {
				logger.Errorf("Abort with error unauthorized err: %v", errAbort)
			}
			return
		}
	}
}

func findAuthorizationToken(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
