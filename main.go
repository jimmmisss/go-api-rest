package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/routes"
)

func main() {
	database.Conexao()
	fmt.Println("Iniciando o servidor rest")
	routes.HandleRequest()
}
