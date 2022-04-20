package main

import (
	"improved_potato/model"
	"improved_potato/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
