package handler

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func AddSSE(r *gin.Engine) {
	seeGroup := r.Group("/api/sse")
	seeGroup.GET("", sse)
}

func sse(ctx *gin.Context) {
	ticker := time.NewTicker(1 * time.Second)
	ctx.Stream(func(w io.Writer) bool {
		select {
		case <-ticker.C:
			ctx.SSEvent("", "heartbeat: "+time.Now().String())
		}
		return true
	})
}
