package models

import "github.com/CRUD/db"

//Produto is
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

//BuscaProdutos is
func BuscaProdutos() []Produto {
	db := db.ConexaoDB() // permite a conexão com o banco

	selectProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC") // cria query que busca os produtos
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() { // pula para próxima linha
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade) // analisa as linhas
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p) // adiciona na lista de produtos
	}

	defer db.Close() // fecha conexão com o banco
	return produtos
}

// CriaNovoProduto is
func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConexaoDB()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)") // preparamos o banco para receber os dados
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade) // executa a inserção no banco trazendo os dados lá da nossa controller
	defer db.Close()
}

// DeletaProduto is
func DeletaProduto(id string) {
	db := db.ConexaoDB()

	deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)
	defer db.Close()
}

// EditaProduto is
func EditaProduto(id string) Produto {
	db := db.ConexaoDB()

	produtoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoEditar := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoEditar.Id = id
		produtoEditar.Nome = nome
		produtoEditar.Descricao = descricao
		produtoEditar.Quantidade = quantidade
		produtoEditar.Preco = preco
	}

	defer db.Close()
	return produtoEditar
}

// AtualizaProduto is
func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConexaoDB()

	AtualizaProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
