package main

import (
	"example.com/gin_realworld/cache"
	"example.com/gin_realworld/server"
)

func main() {
	cache.InitRedis()
	server.RunHTTPServer()
}
