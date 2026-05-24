package handler

import (
	"net/http"

	"example.com/gin_realworld/storage"
	"github.com/gin-gonic/gin"
)

func AddTagsHandler(r *gin.Engine) {
	tagsGroup := r.Group("api/tags")
	tagsGroup.GET("", listPopularTags)
}

func listPopularTags(ctx *gin.Context) {
	tags, err := storage.ListPopularTags(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"tags": tags})
}
