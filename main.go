package main

import (
	"net/http"

	"github.com/CRUD/routes"
)

func main() {
	routes.CarregaRotas()             // permite que alternemos entre as páginas
	http.ListenAndServe(":8000", nil) // porta que estamos rodando nossa aplicação
}
