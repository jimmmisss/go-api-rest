package routes

import (
	"go-api-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/personalidades", controllers.Personalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
