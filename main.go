package main

import (
	"inwin/routes"
)

func main() {
	router := router.SetupRouter()
	router.Run(":8888")
}