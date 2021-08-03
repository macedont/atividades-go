package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //driver de conexao com o mysql
)

//Conectar abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	conexao := "macedo:ProjetoGolang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", conexao)

	if erro != nil {
		return nil, erro
	}

	if db.Ping() != nil { //ping verifica se o banco realmente está conectado
		return nil, erro
	}

	return db, nil
}
