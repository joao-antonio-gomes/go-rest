package routes

import (
	"go-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personas", controllers.FindAllPersonas).Methods("GET")
	r.HandleFunc("/api/personas/{id}", controllers.FindPersonaById).Methods("GET")
	r.HandleFunc("/api/personas", controllers.CreatePersona).Methods("POST")
	r.HandleFunc("/api/personas/{id}", controllers.DeletePersona).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
