package server

import (
	"time"

	"example.com/gin_realworld/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunHTTPServer() {
	r := gin.Default()
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:4273"}                   // 允许的来源
	//config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 允许的方法
	//config.AllowHeaders = []string{"Content-Type", "Authorization"}
	//r.Use(cors.New(config))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	handler.AddUserHandler(r)
	handler.AddArticlesHandler(r)
	handler.AddTagsHandler(r)
	handler.AddArticleCommentHandler(r)
	handler.AddSSE(r)
	handler.AddWebsocket(r)
	r.Static("/images", "./images")
	r.Run()
}
