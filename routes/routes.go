package routes

import (
	"github.com/gorilla/mux"
	"go-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personas", controllers.AllPersonas)
	log.Fatal(http.ListenAndServe(":8000", r))
}
