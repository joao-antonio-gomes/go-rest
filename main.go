package main

import (
	"go-rest/database"
	"go-rest/models"
	"go-rest/routes"
	"log"
)

func main() {
	models.Personas = []models.Persona{
		{Name: "Name 1", Biography: "Biography 1"},
		{Name: "Name 2", Biography: "Biography 2"},
	}
	database.ConnectDatabase()
	log.Println("Beginning REST server with Golang http://localhost:8000")
	routes.HandleRequest()
}
