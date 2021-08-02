package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "macedo:ProjetoGolang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatalln(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatalln(erro)
	}

	fmt.Println("A conexão está aberta")

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatalln(erro)
	}

	defer linhas.Close()
	fmt.Println(linhas)
}