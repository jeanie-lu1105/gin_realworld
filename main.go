package main

import (
	"example.com/gin_realworld/config"
	"example.com/gin_realworld/server"
)

func main() {
	config.GetSecret()
	server.RunHTTPServer()
}
