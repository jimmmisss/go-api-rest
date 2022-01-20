package routes

import (
	"github.com/gorilla/mux"
	"go-api-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.BuscaTodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaNovaPersonalidade).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", r))
}
