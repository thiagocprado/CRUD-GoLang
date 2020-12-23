package routes

import (
	"net/http"

	"github.com/CRUD/controllers"
)

// CarregaRotas is
func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)        // função que renderiza a página inicial
	http.HandleFunc("/new", controllers.New)       // função que renderiza a parte de cadastro
	http.HandleFunc("/insert", controllers.Insert) // função que faz a inserção no banco de dados
	http.HandleFunc("/delete", controllers.Delete) // função que faz a remoção do banco de dados
	http.HandleFunc("/edit", controllers.Edit)     // função que faz a edição do produto
	http.HandleFunc("/update", controllers.Update) // função que faz o update no banco de dados
}
