package main

import (
	"github.com/andersongni/gin-api-rest/database"
	"github.com/andersongni/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
