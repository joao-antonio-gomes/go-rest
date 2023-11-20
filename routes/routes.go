package routes

import (
	"go-rest/controllers"
	"go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personas", controllers.FindAllPersonas).Methods("GET")
	r.HandleFunc("/api/personas/{id}", controllers.FindPersonaById).Methods("GET")
	r.HandleFunc("/api/personas", controllers.CreatePersona).Methods("POST")
	r.HandleFunc("/api/personas/{id}", controllers.DeletePersona).Methods("DELETE")
	r.HandleFunc("/api/personas/{id}", controllers.EditPersona).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8000", r))
}
