package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/CRUD/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) // encapsula nossos templates no caso, as nossas páginas web

// Index is
func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaProdutos()         // trás nossos produtos cadastros
	temp.ExecuteTemplate(w, "Index", todosOsProdutos) // renderiza a página
}

// New is
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil) // renderiza a página
}

// Insert is
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // se o método da requisição for POST
		nome := r.FormValue("nome") // método que utilizamos para pegar os valores do formulário
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64) // convertemos a string para Float
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade) // convertemos a string para Int
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida) // envia os dados para a model para adiconarmos o produto
	}
	http.Redirect(w, r, "/", 301) // retorna para a página principal
}

// Delete is
func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id") // conseguimos pegar o id que está na URL
	models.DeletaProduto(idDoProduto)      // nos conecta com a função do banco que deleta o produto, não precisamos converter pra int
	http.Redirect(w, r, "/", 301)
}

// Edit is
func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

// Update is
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		precoConvertidoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.AtualizaProduto(idConvertido, nome, descricao, precoConvertidoFloat, quantidadeConvertidaInt) // envia os dados para a model para adiconarmos o produto
	}
	http.Redirect(w, r, "/", 301)
}
