package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

//CriarUsuario é a função que cria usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	requisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("falha na requisição"))
		return
	}

	var usr usuario

	if erro = json.Unmarshal(requisicao, &usr); erro != nil {
		w.Write([]byte("Houve um erro ao converter o usuário a um struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Houve um erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) values(?, ?)")
	if erro != nil {
		w.Write([]byte("Houve um erro ao fazer o statement"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(usr.Nome, usr.Email)
	if erro != nil {
		w.Write([]byte("Houve um ero ao executar a inserção"))
		return
	}

	idInsert, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Houve um ero ao recuperar o id da inserção"))
		return
	}

	//STATUS CODES
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInsert)))
}

//BuscarUsuarios retorna todos os usuários salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Houve um erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		w.Write([]byte("Houve um erro ao realizar a consulta com o banco"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Houve um erro ao scannear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Houve um erro ao converter os usuários para json"))
		return
	}
}

//BuscarUsuario retorna um usuário salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	ID, erro := strconv.ParseInt(param["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Houve um erro ao converter o parametro em inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Houve um erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("SELECT * FROM usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Houve um erro ao executar fazer a consulta"))
		return
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Houve um erro ao scannear o usuário"))
			return
		}
	}
	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Nenhum usuário encontrado"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Houve um erro ao colocar o usuário em json"))
		return
	}
}

//AtualizarUsuario é a função que atualiza os dados do usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoResquest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Houve um ero ao pegar o corpo da requisição"))
		return
	}

	var usr usuario

	if erro := json.Unmarshal(corpoResquest, &usr); erro != nil {
		w.Write([]byte("Houve um ero ao passar os dados de json para struct"))
		return
	}

	if usr.Nome == "" || usr.Email == "" {
		w.Write([]byte("Atenção você não pode enviar campos nulos!"))
		return
	}

	param := mux.Vars(r)
	ID, erro := strconv.ParseInt(param["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Houve um ero ao pegar o parametro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Houve um erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Houve um erro ao executar a query de update"))
		return
	}
	defer statement.Close()

	if _, erro = statement.Exec(usr.Nome, usr.Email, ID); erro != nil {
		w.Write([]byte("Houve um erro ao executar o statement"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuário atualizado com sucesso."))
}

//DeleteUsuario é a função que deleta um usuário do banco
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	ID, erro := strconv.ParseInt(param["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Houve um erro ao converter o parametro."))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Houve um erro ao conectar com banco."))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Houve um erro ao realizar o statement."))
		return
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		w.Write([]byte("Houve um erro ao realizar o delete."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuário deletado com sucesso."))
}
