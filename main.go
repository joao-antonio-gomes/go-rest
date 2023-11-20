package main

import (
	"go-rest/routes"
	"log"
)

func main() {
	log.Println("Beginning REST server with Golang http://localhost:8000")
	routes.HandleRequest()
}
