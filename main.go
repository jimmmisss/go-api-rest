package main

import (
	"fmt"
	"go-api-rest/models"
	"go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{"Wesley Pereira", "História complicada"},
		{"Fadia Pereira", "História de boa"},
	}

	fmt.Println("Iniciando o servidor rest")
	routes.HandleRequest()
}
