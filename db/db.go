package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// ConexaoDB is
func ConexaoDB() *sql.DB {
	conexao := "user=postgres dbname=thigas_loja password=thiago123 host=localhost sslmode=disable" // faz a conexão com o banco
	db, err := sql.Open("postgres", conexao)                                                        // abre o banco
	if err != nil {
		panic(err.Error())
	}
	return db
}
