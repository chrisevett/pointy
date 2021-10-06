package main

import (
	"chrisevett/pointy/v1/routers"
	"chrisevett/pointy/v1/services"
)

func main() {
	services.ConnectDatabase()
	routers.Run()
}
