package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest/database"
	"go-rest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func FindAllPersonas(w http.ResponseWriter, r *http.Request) {
	var p []models.Persona

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func FindPersonaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var persona models.Persona
	database.DB.First(&persona, id)

	json.NewEncoder(w).Encode(persona)
}
