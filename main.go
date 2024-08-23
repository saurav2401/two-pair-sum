package main

import (
	"backend/router"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Go server going to start....")
	route := router.SetupRouter()

	err := route.Run(":8081")
	if err != nil {
		log.Fatal("server failed to start :", err)
	}
}
