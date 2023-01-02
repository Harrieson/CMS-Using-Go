package main

import (
	"Go-cms/model"
	"Go-cms/server"
)

func main() {
	model.SetUp()
	server.SetUpAndListen()
}
