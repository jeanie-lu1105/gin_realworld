package server

import (
	"example.com/gin_realworld/handler"
	"github.com/gin-gonic/gin"
)

func RunHTTPServer() {
	r := gin.Default()
	handler.AddUserHandler(r)
	handler.AddArticlesHandler(r)
	handler.AddTagsHandler(r)
	r.Run()
}
