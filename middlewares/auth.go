package middlewares

import (
	"net/http"
	"strings"

	"example.com/gin_realworld/logger"
	"example.com/gin_realworld/security"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	log := logger.New(ctx)
	token := ctx.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	claims, ok, err := security.VerifyJWT(token)
	if err != nil || !ok {
		log.WithError(err).Info("Error verifying JWT")
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	ctx.Set("user", claims)
	ctx.Next()
}

func AuthMiddlewareCookie(ctx *gin.Context) {
	log := logger.New(ctx)
	token, err := ctx.Cookie("token")
	if err == http.ErrNoCookie {
		token = ctx.GetHeader("Authorization")
	}

	token = strings.TrimPrefix(token, "Token ")
	token = strings.TrimPrefix(token, "Bearer ")
	claims, ok, err := security.VerifyJWT(token)
	if err != nil || !ok {
		log.WithError(err).Info("Error verifying JWT")
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	ctx.Set("user", claims)
	ctx.Next()
}
