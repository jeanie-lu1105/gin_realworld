package main

import (
	"example.com/gin_realworld/config"
	"example.com/gin_realworld/server"
	"example.com/gin_realworld/storage"
)

func main() {
	config.GetSecret()
	storage.AddUsers()
	server.RunHTTPServer()
}
