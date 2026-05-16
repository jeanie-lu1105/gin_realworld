package handler

import (
	"example.com/gin_realworld/logger"
	"github.com/gin-gonic/gin"
)

func TagsHandler(r *gin.Engine) {
	usersGroup := r.Group("/api/tags")
	usersGroup.GET("", listTags)
}

func listTags(ctx *gin.Context) {
	log := logger.New(ctx)
	log.WithContext(ctx).Info("list tags api called")
	return
}
