package main

import (
	"github.com/cyctw/go-crud/models"
	"github.com/cyctw/go-crud/server"
)

func main() {
	models.ConnectDatabase()
	server.Init()
}
