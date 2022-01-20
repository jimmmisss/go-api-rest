package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-rest/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func Personalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
